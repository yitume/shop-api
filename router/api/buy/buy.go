package buy

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/yitume/shop-api/router/mdw"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/milkbobo/gopay"
	"github.com/milkbobo/gopay/client"
	"github.com/milkbobo/gopay/common"
	"github.com/milkbobo/gopay/constant"
	"github.com/pkg/errors"
	"github.com/satori/uuid"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/resp"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/model/view"
	"github.com/yitume/shop-api/pkg/logic"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/router/mdw/wechat"
	"github.com/yitume/shop-api/service"
)

type goodsSkuInfo struct {
	Id      int
	Stock   int
	SaleNum int
}

/**
  goods_amount;
  pay_amount;
  goods_freight_list;
  freight_unified_fee;
  freight_template_fee;
  pay_freight_fee;
  subtotal;
*/

func Calculate(c *gin.Context) {
	req := trans.ReqBuyCalculate{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)

	addressInfo, err := service.Address.InfoX(c, mysql.Conds{
		"id":  req.AddressId,
		"uid": uid,
	})
	if err != nil && err == gorm.ErrRecordNotFound {
		err = errors.New("收货地址没找到")
		return
	}

	list := service.Cart.ListArr(c, uid, req.CartIds)

	output, err := calculate(addressInfo, list)
	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	base.JSON(c, base.MsgOk, output)
}

/**
 * 创建订单
 */
func Create(c *gin.Context) {
	req := trans.ReqBuyCreate{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request error")
		return
	}
	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)

	addressInfo, err := service.Address.InfoX(c, mysql.Conds{
		"id":      req.AddressId,
		"uid":     uid,
		"open_id": openId,
	})
	if err != nil && err == gorm.ErrRecordNotFound {
		err = errors.New("收货地址没找到")
		return
	}

	list := service.Cart.ListArr(c, uid, req.CartIds)

	tx := model.Db.Begin()

	create := &mysql.OrderPay{
		PaySn:    strings.ReplaceAll(uuid.NewV4().String(), "-", ""),
		UserId:   uid,
		PayState: 0,
	}

	err = service.OrderPay.Create(c, tx, create)
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr, "request error")
		return
	}
	calculateValue, err := calculate(addressInfo, list)

	orderCreate := &mysql.Order{
		Id:                 0,
		Sn:                 strings.ReplaceAll(uuid.NewV4().String(), "-", ""),
		PaySn:              create.PaySn,
		Uid:                uid,
		UserName:           wechat.WechatUserName(c),
		UserPhone:          addressInfo.MobilePhone,
		UserEmail:          "",
		CreateTime:         time.Now().Unix(),
		PaymentCode:        "online",
		PayName:            "online",
		PaymentTime:        0,
		FinnshedTime:       0,
		GoodsAmount:        calculateValue.GoodsAmount,
		GoodsNum:           calculateValue.GoodsNum,
		Amount:             calculateValue.PayAmount,
		PdAmount:           0,
		FreightFee:         calculateValue.PayFreightFee,
		FreightUnifiedFee:  calculateValue.FreightUnifiedFee,
		FreightTemplateFee: calculateValue.FreightTemplateFee,
		State:              service.OrderStateNew,
		RefundAmount:       0,
		RefundState:        0,
		LockState:          0,
		DelayTime:          0,
		Points:             0,
		EvaluateState:      0,
		IsPrint:            0,
		TradeNo:            "",
		OutRequestNo:       "",
		WechatOpenid:       "",
		From:               1,
		DeleteTime:         0,
		AllAgreeRefound:    0,
		PayableTime:        time.Now().Unix() + 24*3600, // 订单什么时间过期，先写24h
		GoodsType:          0,
		GroupId:            0,
		GroupIdentity:      0,
		GroupSign:          "",
		GroupPeopleNum:     0,
		GroupMenNum:        0,
		GroupState:         0,
		GroupStartTime:     0,
		GroupEndTime:       0,
		GoodsGroupAmount:   0,
		GroupFailTime:      0,
		ReviseAmount:       0,
		ReviseFreightFee:   0,
		OpenId:             openId,
	}

	err = service.Order.Create(c, model.Db, orderCreate)
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	// 拓展订单表创建
	orderExtendCreate := &mysql.OrderExtend{
		Id: orderCreate.Id,
		ReciverInfo: mysql.OrderExtendReciverInfoJson{
			Name:          addressInfo.Truename,
			CombineDetail: addressInfo.CombineDetail,
			Phone:         addressInfo.MobilePhone,
			Type:          addressInfo.Type,
			Address:       addressInfo.Address,
		},
		ReciverName:       addressInfo.Truename,
		ReceiverPhone:     addressInfo.MobilePhone,
		ReciverProvinceId: addressInfo.ProvinceId,
		ReciverCityId:     addressInfo.CityId,
		ReciverAreaId:     addressInfo.AreaId,
	}

	err = service.OrderExtend.Create(c, model.Db, orderExtendCreate)
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	cartIds := make([]int, 0)
	for _, value := range list {
		cartIds = append(cartIds, value.Id)
		create4 := &mysql.OrderGoods{
			OrderId:         orderExtendCreate.Id,
			GoodsId:         value.GoodsId,
			GoodsSkuId:      value.GoodsSkuId,
			GoodsTitle:      value.GoodsTitle,
			GoodsSpec:       mysql.OrderGoodsGoodsSpecJson(value.GoodsSpec),
			GoodsPrice:      value.GoodsPrice,
			GoodsPayPrice:   value.GoodsPrice * float64(value.GoodsNum),
			GoodsNum:        value.GoodsNum,
			GoodsImg:        value.GoodsSkuImg,
			GoodsFreightFee: value.GoodsFreightFee,
			GoodsFreightWay: value.GetFreightWay(),
			UserId:          uid,
			CreateTime:      time.Now().Unix(),
			GoodsType:       1,
		}

		err = service.OrderGoods.Create(c, tx, create4)
		if err != nil {
			tx.Rollback()
			base.JSON(c, base.MsgErr)
			return
		}
	}

	createLog := &mysql.OrderLog{
		OrderId:    orderCreate.Id,
		Msg:        "买家下单",
		Role:       "buyer",
		OrderState: service.OrderStateNew,
	}

	err = service.OrderLog.Create(c, tx, createLog)
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}
	// todo
	err = updateGoodsStorageNum(tx, list)
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}
	// todo 优化,delete by in
	for _, cartId := range cartIds {
		err = service.Cart.DeleteX(c, tx, mysql.Conds{
			"id":      cartId,
			"open_id": openId,
		})

		if err != nil {
			tx.Rollback()
			base.JSON(c, base.MsgErr)
			return
		}
	}

	tx.Commit()
	base.JSON(c, base.MsgOk)
}

func calculate(addressInfo mysql.Address, list []logic.CartItem) (output resp.BuyCalculate, err error) {
	freightList := make([]view.Freight, 0)

	// 根据运费计算规则一
	var freightUnifiedFee float64
	var freightTemplateFee float64
	var goodsAmount float64
	var payFreightFee float64
	var goodsNum int

	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	base.JSON(c, base.MsgErr, "数据库存在异常")
	// 	return
	// }

	for _, value := range list {
		if value.GoodsNum > value.GoodsStock {
			err = errors.New("库存不够")
			return
		}
		valuePointer := &value
		var cost float64
		cost, err = valuePointer.FreightFeeByAddress(addressInfo)
		if err != nil {
			err = errors.New("数据异常")
			return
		}

		freightWay := valuePointer.GetFreightWay()

		freightList = append(freightList, view.Freight{
			GoodsSkuId: valuePointer.GoodsSkuId,
			FreightFee: cost,
			FreightWay: freightWay,
		})

		if freightWay == "goods_freight_unified" {
			// todo 为什么是=
			freightUnifiedFee = valuePointer.GoodsFreightFee
		} else {
			freightTemplateFee += valuePointer.GoodsFreightFee
		}

		goodsAmount += valuePointer.GoodsPrice * float64(valuePointer.GoodsNum)
		goodsNum += valuePointer.GoodsNum
	}

	payFreightFee = freightUnifiedFee + freightTemplateFee

	output = resp.BuyCalculate{
		GoodsAmount:        goodsAmount,
		PayAmount:          goodsAmount + payFreightFee,
		GoodsFreightList:   freightList,
		FreightUnifiedFee:  freightUnifiedFee,
		FreightTemplateFee: freightTemplateFee,
		PayFreightFee:      payFreightFee,
		SubTotal:           goodsAmount + payFreightFee,
		GoodsNum:           goodsNum,
	}
	return
}

func updateGoodsStorageNum(tx *gorm.DB, list []logic.CartItem) (err error) {

	goodsSkuInfos := make(map[int]goodsSkuInfo, 0)

	for _, value := range list {
		// 同一款产品的sku数据相加更新到主表
		goodsId := value.GoodsId
		if tmpValue, ok := goodsSkuInfos[goodsId]; !ok {
			goodsSkuInfos[goodsId] = goodsSkuInfo{
				Id:      value.GoodsId,
				Stock:   0,
				SaleNum: value.GoodsNum,
			}
		} else {
			tmpValue.SaleNum += value.GoodsNum
			goodsSkuInfos[goodsId] = tmpValue
		}

		err = tx.Exec("UPDATE goods_sku set `stock`=stock-? , `sale_num`=sale_num+? WHERE id=?", value.GoodsNum, value.GoodsNum, goodsId).Error
		if err != nil {
			return
		}
	}

	for _, value := range goodsSkuInfos {
		err = tx.Exec("UPDATE goods_sku set `stock`=stock-? , `sale_num`=sale_num+? WHERE id=?", value.SaleNum, value.SaleNum, value.Id).Error
		if err != nil {
			return
		}
	}
	return
}

type WechatPayConfig struct {
	AppID          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
	MiniAppID      string `json:"mini_app_id"`
	MiniAppSecret  string `json:"mini_app_secret"`
	APPID          string `json:"appid"`
	MchID          string `json:"mch_id"`
	Key            string `json:"key"`
	ApiclientCert  string `json:"apiclient_cert"`
	ApiclientKey   string `json:"apiclient_key"`
	CallbackDomain string `json:"callback_domain"`
}

// Pay 订单支付接口
func Pay(c *gin.Context) {
	req := trans.ReqBuyPay{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request error")
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	domain := mdw.DomainName(c)

	setting, err := service.Setting.InfoX(c, mysql.Conds{
		"key":    req.PaymentCode,
		"status": 1,
	})
	if err != nil {
		err = errors.New("系统暂时不支持选定的支付方式")
		return
	}

	order, err := service.Order.InfoX(c, mysql.Conds{
		"pay_sn":  req.PaySn,
		"uid":     uid,
		"state":   service.OrderStateNew,
		"open_id": openId,
	})
	if err != nil && err == gorm.ErrRecordNotFound {
		err = errors.New("该订单不存在")
		return
	}

	orderPay, err := service.OrderPay.InfoX(c, mysql.Conds{
		"pay_sn":    req.PaySn,
		"pay_state": service.OrderPayStateNew,
	})
	if err != nil && err == gorm.ErrRecordNotFound {
		err = errors.New("该订单不存在")
		return
	}

	payAmount := order.Amount
	if order.ReviseAmount > 0 {
		payAmount = order.ReviseAmount
	}

	result := map[string]string{}
	switch req.PaymentChannel {
	case "wechat_mini":
		userOpen, err := service.UserOpen.InfoX(c, mysql.Conds{"uid": uid, "genre": 1})
		if err != nil {
			err = errors.New("暂未获取到微信关联用户信息")
			return
		}

		config := setting.Config.Data.(map[string]interface{})
		client.InitWxMiniProgramClient(&client.WechatMiniProgramClient{
			AppID:      config["mini_app_id"].(string),
			MchID:      config["mch_id"].(string),
			Key:        config["key"].(string),
			PrivateKey: nil,
			PublicKey:  nil,
		})

		charge := new(common.Charge)
		charge.PayMethod = constant.WECHAT_MINI_PROGRAM
		charge.MoneyFee = payAmount
		charge.TradeNum = req.PaySn
		charge.Describe = fmt.Sprintf("商品购买_%s", orderPay.PaySn)
		charge.TradeNum = orderPay.PaySn
		charge.CallbackURL = fmt.Sprintf("%s/api/%s/callback/wechatMiniNotify", config["callback_domain"].(string), domain)
		charge.OpenID = userOpen.MiniOpenid
		fdata, err := gopay.Pay(charge)
		fmt.Println("charge.CallbackURL------>", charge.CallbackURL)
		if err != nil {
			fmt.Println("=============Pay error===============")
			fmt.Println(config["mini_app_id"].(string))
			fmt.Println(config["mch_id"].(string))
			fmt.Println(charge.OpenID)
			fmt.Println(err)
			fmt.Println("=============Pay error===============")
			base.JSON(c, base.MsgErr)
			return
		}

		result = fdata
		break
	default:
		break
	}

	base.JSON(c, base.MsgOk, result)
	return
}

func WechatMiniNotify(c *gin.Context) {
	domain := mdw.DomainName(c)

	var openUser mysql.Biz
	model.Db.Select("open_id,domain").Where("domain=?", domain).Find(&openUser)
	fmt.Println(openUser)

	setting, err := service.Setting.InfoX(c, mysql.Conds{
		"key":     "wechat",
		"open_id": openUser.OpenId,
	})
	fmt.Println("setting------>", setting)
	if err != nil {
		model.Logger.Error("wechat mini notify error", zap.String("err", err.Error()))
		return
	}
	fmt.Println("setting.Config------>", setting.Config)

	// todo 以后在根据type类型判断为哪种结构体，先实现功能，直接断言

	config := setting.Config.Data.(map[string]interface{})
	client.InitWxAppClient(&client.WechatAppClient{
		AppID:      config["mini_app_id"].(string),
		MchID:      config["mch_id"].(string),
		Key:        config["key"].(string),
		PrivateKey: nil,
		PublicKey:  nil,
	})

	//返回支付结果
	wechatResult, err := gopay.WeChatAppCallback(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 检查返回值
	if wechatResult.ResultCode != "SUCCESS" {
		fmt.Println(err)
		return
	}

	//接下来处理自己的逻辑
	fmt.Println(wechatResult)
	err = service.Order.Pay(c, wechatResult.OutTradeNO, "wechat", wechatResult.TransactionID, openUser.OpenId)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
