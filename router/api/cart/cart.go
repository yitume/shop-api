package cart

import (
	"github.com/yitume/shop-api/pkg/common/code"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/yitume/shop-api/router/mdw/wechat"
	"github.com/yitume/shop-api/service"
)

func Info(c *gin.Context) {
	req := trans.ReqCartInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgCarInfoReqErr)
		return
	}

	resp, err := service.Cart.InfoA(c, wechat.WechatUid(c), req.GoodsSkuId)
	if err != nil {
		base.JSON(c, code.MsgCarInfoEmptyErr)
		return
	}
	base.JSON(c, base.MsgOk, resp)

}

/**
  id;
  goods_sku_id;
  goods_num;
  goods_id;
  goods_title;
  goods_is_on_sale;
  goods_freight_fee;
  goods_freight_id;
  goods_pay_type;
  goods_price;
  goods_spec;
  goods_weight;
  goods_stock;
  goods_sku_img;
  goods_freight_areas;
  create_time;
  checked;
*/

func List(c *gin.Context) {
	uid := wechat.WechatUid(c)
	output := service.Cart.ListArr(c, uid, []int{})
	base.JSON(c, code.MsgOk, output)

}

func Create(c *gin.Context) {
	req := trans.ReqCartCreateInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgCartCreateReqErr)
		return
	}
	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	cartInfo, _ := service.Cart.InfoX(c, mysql.Conds{
		"goods_sku_id": req.GoodsSkuId,
		"uid":          uid,
	})

	if cartInfo.Id > 0 {
		base.JSON(c, code.MsgCartCreateCartExistErr)
		return
	}

	goodsSkuInfo, _ := service.GoodsSku.InfoX(c, mysql.Conds{
		"id": req.GoodsSkuId,
	})

	if goodsSkuInfo.Id == 0 {
		base.JSON(c, base.MsgErr)
		return
	}
	goodsInfo, _ := service.Goods.InfoX(c, mysql.Conds{
		"id":         goodsSkuInfo.GoodsId,
		"open_id":    openId,
		"is_on_sale": 1,
	})
	if goodsInfo.Id == 0 {
		base.JSON(c, base.MsgErr, "产品不存在")
		return
	}

	// 购物车里原基础增加数量
	quantity := req.Quantity
	if cartInfo.GoodsNum > 0 {
		quantity += cartInfo.GoodsNum
	}

	if goodsSkuInfo.Stock < 1 {
		base.JSON(c, base.MsgErr, "商品库存不足")
		return
	}

	if goodsSkuInfo.Stock < quantity {
		base.JSON(c, base.MsgErr, "库存不足")
		return
	}

	err := service.Cart.Create(c, model.Db, &mysql.Cart{0, uid, req.GoodsSkuId, quantity, 1, time.Now().Unix(), 0, openId})

	if err != nil {
		base.JSON(c, base.MsgErr, "插入失败")
		return
	}
	base.JSON(c, base.MsgOk, "ok")
}

func Update(c *gin.Context) {
	req := trans.ReqCartUpdateInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	cartInfo, _ := service.Cart.InfoX(c, mysql.Conds{
		"goods_sku_id": req.GoodsSkuId,
		"uid":          uid,
		"open_id":      openId,
	})

	if cartInfo.Id == 0 {
		base.JSON(c, base.MsgErr, "购物车不存在")
		return
	}

	goodsSkuInfo, _ := service.GoodsSku.InfoX(c, mysql.Conds{
		"id":      req.GoodsSkuId,
		"open_id": openId,
	})

	if goodsSkuInfo.Id == 0 {
		base.JSON(c, base.MsgErr, "产品不存在")
		return
	}

	goodsInfo, _ := service.Goods.InfoX(c, mysql.Conds{
		"id":         goodsSkuInfo.GoodsId,
		"open_id":    openId,
		"is_on_sale": 1,
	})
	if goodsInfo.Id == 0 {
		base.JSON(c, base.MsgErr, "产品不存在")
		return
	}

	// 库存小于传过来的参数
	if goodsSkuInfo.Stock == 0 {
		base.JSON(c, base.MsgErr, "库存不足")
		return
	}

	if goodsSkuInfo.Stock < req.Quantity {
		err := service.Cart.UpdateX(c, model.Db, mysql.Conds{
			"id":      cartInfo.Id,
			"uid":     uid,
			"open_id": openId,
		}, mysql.Ups{
			"goods_num": goodsSkuInfo.Stock,
		})
		if err != nil {
			base.JSON(c, base.MsgErr, "更新失败")
			return
		}
		base.JSON(c, base.MsgErr, "库存较少，更新给用户")
		return
	}

	err := service.Cart.UpdateX(c, model.Db, mysql.Conds{
		"id":      cartInfo.Id,
		"uid":     uid,
		"open_id": openId,
	}, mysql.Ups{
		"goods_num":    req.Quantity,
		"goods_sku_id": goodsSkuInfo.Id,
	})
	if err != nil {
		base.JSON(c, base.MsgErr, "更新失败")
		return
	}
	base.JSON(c, base.MsgOk)

}

func Del(c *gin.Context) {
	req := trans.ReqCartDel{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	err := service.Cart.DeleteX(c, model.Db, mysql.Conds{
		"uid":          uid,
		"open_id":      openId,
		"goods_sku_id": mysql.Cond{"in", req.GoodsSkuIds},
	})

	if err != nil {
		base.JSON(c, base.MsgErr, "删除失败")
		return
	}
	base.JSON(c, base.MsgOk, "删除成功")

}
func Check(c *gin.Context) {
	req := trans.ReqCartCheck{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)

	err := service.Cart.UpdateX(c, model.Db, mysql.Conds{
		"uid":          uid,
		"open_id":      openId,
		"goods_sku_id": mysql.Cond{"in", req.GoodsSkuIds},
	}, mysql.Ups{
		"is_check": req.IsCheck,
	})

	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	base.JSON(c, base.MsgOk)
}

func Exist(c *gin.Context) {
	req := trans.ReqCartInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	_, err := service.Cart.InfoX(c, mysql.Conds{
		"uid":          wechat.WechatUid(c),
		"goods_sku_id": req.GoodsSkuId,
		"open_id":      mdw.OpenId(c),
	})

	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	base.JSON(c, base.MsgOk)

}

func TotalNum(c *gin.Context) {
	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	base.JSON(c, base.MsgOk, map[string]interface{}{
		"total_num": service.Cart.TotalNum(uid, openId),
	})

}
