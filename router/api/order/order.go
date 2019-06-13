package order

import (
	"fmt"

	"github.com/yitume/shop-api/router/mdw"

	"github.com/yitume/shop-api/pkg/common"
	"github.com/yitume/shop-api/router/mdw/wechat"

	"github.com/gin-gonic/gin"

	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/service"
)

func StateNum(c *gin.Context) {
	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	base.JSON(c, base.MsgOk, map[string]interface{}{
		"state_new":       service.Order.GetStateCnt(uid, service.OrderStateNew, openId),
		"state_send":      service.Order.GetStateCnt(uid, service.OrderStateSend, openId),
		"state_success":   service.Order.GetStateCnt(uid, service.OrderStateSuccess, openId),
		"state_close":     service.Order.GetStateCnt(uid, service.OrderStateClose, openId),
		"tate_unevaluate": 0,
		"tate_refund":     0,
	})

}

func Info(c *gin.Context) {
	req := trans.ReqOrderInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	uid := wechat.WechatUid(c)

	orderInfo, err := service.Order.GetOrderInfo(c, req.Id, uid)

	if err != nil {
		base.JSON(c, base.MsgErr, "info not exist")
		return
	}

	base.JSON(c, base.MsgOk, orderInfo)
}
func Cancel(c *gin.Context) {
	req := trans.ReqOrderCancel{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}
	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)

	orderInfo, err := service.Order.InfoByIdAndUid(req.Id, uid, openId)
	if err != nil {
		base.JSON(c, base.MsgErr, "info not exist")
		return
	}
	fmt.Println(orderInfo)

}

/**
id;
    sn;
    pay_sn;
    create_time;
    payment_code;
    pay_name;
    payment_time;
    finnshed_time;
    goods_amount;
    goods_num;
    amount;
    pd_amount;
    freight_fee;
    freight_unified_fee;
    freight_template_fee;
    state;
    refund_amount;
    refund_state;
    lock_state;
    delay_time;
    tracking_no;
    evaluate_state;
    trade_no;
    state_desc;
    payment_name;
    extend_order_extend;
    extend_order_goods;

    if_cancel;
    if_pay;
    if_evaluate;
    if_receive;
*/

func List(c *gin.Context) {
	req := trans.ReqOrderListInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request is error")
		return
	}
	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	conds := mysql.Conds{
		"uid":     uid,
		"open_id": openId,
	}

	// 订单状态查询
	if req.StateType != "" {
		state, ok := common.OrderStates[req.StateType]
		if !ok {
			base.JSON(c, base.MsgErr, "request app list params is error")
			return
		}
		conds["state"] = state
	}

	total, orderMap := service.Order.GetOrderFullsPage(c, trans.ReqPage{
		Current:  req.Page - 1,
		PageSize: req.Rows,
	}, conds, map[string]bool{
		"order_goods":  true,
		"order_extend": true,
	})

	for _, value := range orderMap {
		// 显示取消订单
		value.IsCancel = service.Order.GetOrderOperateState("user_cancel", value.Order)
		// 显示是否需能支付（todo 计算后台过期时间）
		value.IsPay = service.Order.GetOrderOperateState("user_pay", value.Order)
		// 显示退款取消订单
		value.IsRefundCancel = service.Order.GetOrderOperateState("refund_cancel", value.Order)
		// 显示投诉
		value.IsComplain = service.Order.GetOrderOperateState("complain", value.Order)
		// 显示收货
		value.IsReceive = service.Order.GetOrderOperateState("receive", value.Order)
		// 显示锁定中
		value.IsLock = service.Order.GetOrderOperateState("lock", value.Order)
		// 显示物流跟踪
		value.IsDeliver = service.Order.GetOrderOperateState("deliver", value.Order)
		// 显示评价
		value.IsEvaluate = service.Order.GetOrderOperateState("evaluate", value.Order)
	}

	output := make([]service.OrderFull, 0)
	for _, value := range orderMap {
		output = append(output, *value)
	}
	// todo要按时间排序		Sort:     "create_time desc",

	base.JSONList(c, output, total)
}

func ConfirmReceipt(c *gin.Context) {
	req := trans.ReqOrderConfirmReceipt{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}
	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)

	orderInfo, err := service.Order.InfoByIdAndUid(req.Id, uid, openId)
	if err != nil {
		base.JSON(c, base.MsgErr, "info not exist")
		return
	}

	flag := service.Order.UserChangeState(c, "order_receive", orderInfo, uid, wechat.WechatUserName(c), req.StateRemark)
	if !flag {
		base.JSON(c, base.MsgErr, "change state error")
		return
	}
	base.JSON(c, base.MsgOk, "change state ok")
	return

}

func Logistics(c *gin.Context) {
	req := trans.ReqOrderLogistics{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	_, err := service.Order.InfoX(c, mysql.Conds{
		"id":    req.Id,
		"state": mysql.Cond{">", 30},
		"uid":   wechat.WechatUid(c),
	})
	if err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	orderExtendInfo, err := service.OrderExtend.InfoX(c, mysql.Conds{
		"id": req.Id,
	})

	if err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	if orderExtendInfo.TrackingNo != "" && orderExtendInfo.ShipperId > 0 && orderExtendInfo.ExpressId > 0 {
		info, err := service.Express.InfoX(c, mysql.Conds{
			"id": orderExtendInfo.ExpressId,
		})
		if err != nil {
			base.JSON(c, base.MsgOk, "ok", map[string]interface{}{
				"url": "",
			})
			return
		}

		base.JSON(c, base.MsgOk, "ok", map[string]interface{}{
			"url": "https://m.kuaidi100.com/index_all.html?type=" + info.Kuaidi100Code + "&postid=" + orderExtendInfo.TrackingNo,
		})
		return
	}
	base.JSON(c, base.MsgOk, "ok", map[string]interface{}{
		"url": "",
	})
}

func GoodsList(c *gin.Context) {
	req := trans.ReqOrderGoodsList{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	cnt, list := service.OrderGoods.ListPage(c, mysql.Conds{
		"order_id": req.Id,
		"uid":      wechat.WechatUid(c),
	}, trans.ReqPage{
		Current:  0,
		PageSize: 10000,
		Sort:     "id asc",
	})
	base.JSONList(c, list, cnt)
}

func GoodsInfo(c *gin.Context) {
	req := trans.ReqOrderGoodsInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	resp, _ := service.OrderGoods.InfoX(c, mysql.Conds{
		"order_id": req.Id,
		"uid":      wechat.WechatUid(c),
	})
	base.JSON(c, base.MsgOk, map[string]interface{}{
		"info": resp,
	})
}
