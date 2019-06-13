package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/model/view"
	"github.com/yitume/shop-api/pkg/common"
)

// 取消订单
const OrderStateClose = 0

// 未支付订单
const OrderStateNew = 10

// 已支付
const OrderStatePay = 20

// 已发货
const OrderStateSend = 30

// 已收货，交易成功
const OrderStateSuccess = 40

const ComplainTimeLimit = 24 * 60 * 60 * 7

const ORDER_EVALUATE_TIME = 1296000 // todo 需要可修改

// 未支付订单 OrderPay
const OrderPayStateNew = 0

// 已支付订单 OrderPay
const OrderPayStatePay = 1

// 待付款
const GroupStateNew = 0

// 正在进行中(待开团)
const GroupStatePay = 1

// 拼团成功
const GroupStateSuccess = 2

// 拼团失败
const GroupStateFail = 3

// 未评价
//const OrderStateUnevaluate  = 50

//const OrderStateRefund  = 60

type OrderFull struct {
	Order             mysql.Order `json:"order"`
	StateDesc         string      `json:"state_desc"`
	GroupStateDesc    string      `json:"group_state_desc"`
	PaymentName       string      `json:"payment_name"`
	mysql.UserOpen    `json:"extend_user"`
	mysql.OrderExtend `json:"extend_order_extend"`
	ExtendOrderGoods  []ExtendOrderGoods `json:"extend_order_goods"`

	// 新增信息
	IsCancel       bool `json:"is_cancel"`
	IsPay          bool `json:"is_pay"`
	IsRefundCancel bool `json:"is_refund_cancel"`
	IsComplain     bool `json:"is_complain"`
	IsReceive      bool `json:"is_receive"`
	IsLock         bool `json:"is_lock"`
	IsDeliver      bool `json:"is_deliver"`
	IsEvaluate     bool `json:"is_evaluate"`
}

type ExtendOrderGoods struct {
	mysql.OrderGoods
	IfRefund bool `json:"if_refund"`
}

func (*order) GetStateCnt(uid int, state int, openId int) (cnt int) {
	model.Db.Model(mysql.Order{}).Where("uid = ? AND state = ? AND open_id = ?", uid, state, openId).Count(&cnt)
	return
}

func (*order) InfoByIdAndUid(id int, uid int, openId int) (resp mysql.Order, err error) {
	err = model.Db.Model(mysql.Order{}).Where("id = ? AND uid = ? AND open_id", id, uid, openId).Find(&resp).Error
	return
}

// 取单条订单信息
func (o *order) GetOrderInfo(c *gin.Context, id int, uid int) (resp trans.RespOrderInfo, err error) {
	var orderInfo mysql.Order
	orderInfo, err = o.InfoX(c, mysql.Conds{
		"id":  id,
		"uid": uid,
	})
	if err != nil {
		return
	}

	_, orderLog := OrderLog.ListPage(c, mysql.Conds{
		"order_id": orderInfo.Id,
	}, trans.ReqPage{
		Current:  0,
		PageSize: 1000,
	})

	extendInfo, _ := OrderExtend.Info(c, orderInfo.Id)

	//追加返回商品信息

	//取商品列表
	extendOrderGoodsOrigin, _ := OrderGoods.List(c, mysql.Conds{"id": orderInfo.Id}, "id desc")

	extendOrderGoods := make([]trans.FaOrderExtend, 0)

	for _, value := range extendOrderGoodsOrigin {
		// 退款平台处理状态 默认0处理中(未处理) 10拒绝(驳回) 20同意 30成功(已完成) 50取消(用户主动撤销) 51取消(用户主动收货)
		// 不可退款
		valueView := trans.FaOrderExtend{}

		if value.RefundId > 0 && funk.ContainsInt([]int{20, 30, 51}, value.RefundHandleState) {
			valueView.IsRefund = 0
		} else {
			valueView.IsRefund = 1
		}

		var refundState int
		//refund_state 0不显示申请退款按钮 1显示申请退款按钮 2显示退款中按钮 3显示退款完成
		if orderInfo.RefundState <= 10 {
			refundState = 0
		}

		if orderInfo.RefundState > 10 && orderInfo.LockState == 0 && value.RefundId == 0 {
			refundState = 1

		}

		if orderInfo.RefundState > 10 && (orderInfo.LockState != 0 || value.RefundId != 0) && value.RefundHandleState == 30 {
			refundState = 3
		}

		if orderInfo.RefundState > 10 && (orderInfo.LockState != 0 || value.RefundId != 0) && value.RefundHandleState != 30 {
			refundState = 2
		}
		valueView.RefundStatus = refundState

		extendOrderGoods = append(extendOrderGoods, valueView)

	}

	orderCondition := view.OrderCondition{}
	// 显示取消订单
	orderCondition.IsCancel = o.GetOrderOperateState("user_cancel", orderInfo)
	// 显示是否需能支付（todo 计算后台过期时间）
	orderCondition.IsPay = o.GetOrderOperateState("user_pay", orderInfo)
	// 显示退款取消订单
	orderCondition.IsRefundCancel = o.GetOrderOperateState("refund_cancel", orderInfo)
	// 显示投诉
	orderCondition.IsComplain = o.GetOrderOperateState("complain", orderInfo)
	// 显示收货
	orderCondition.IsReceive = o.GetOrderOperateState("receive", orderInfo)
	// 显示锁定中
	orderCondition.IsLock = o.GetOrderOperateState("lock", orderInfo)
	// 显示物流跟踪
	orderCondition.IsDeliver = o.GetOrderOperateState("deliver", orderInfo)
	// 显示评价
	orderCondition.IsEvaluate = o.GetOrderOperateState("evaluate", orderInfo)

	resp = trans.RespOrderInfo{
		Info:              orderInfo,
		OrderCondition:    orderCondition,
		OrderLog:          orderLog,
		ExtendOrderExtend: extendInfo,
		ExtendOrderGoods:  extendOrderGoods,
	}
	return
}

//  返回是否允许某些操作
func (*order) GetOrderOperateState(operateState string, orderInfo mysql.Order) bool {
	var state bool
	switch operateState {
	//买家支付
	case "user_pay":
		state = orderInfo.State == OrderStateNew && time.Now().Unix() < orderInfo.PayableTime
	//买家取消订单
	case "user_cancel":
		//$state = ($order_info['state'] == \App\Logic\Order::state_new) || ($order_info['payment_code'] == 'offline' && $order_info['state'] == \App\Logic\Order::state_pay);
		state = orderInfo.State == OrderStateNew || orderInfo.PaymentCode == "offline" && orderInfo.State == OrderStatePay
	//买家取消订单
	case "refund_cancel":
		//$state = $order_info['refund_state'] == 1 && !intval( $order_info['lock_state'] );
		// todo
		state = orderInfo.RefundState == 1 && orderInfo.LockState != 1
		//商家取消订单
	case "cancel":
		state = orderInfo.State == OrderStateNew || orderInfo.PaymentCode == "offline" && funk.ContainsInt([]int{OrderStatePay, OrderStateSend}, orderInfo.State)
		//平台取消订单
	case "system_cancel":
		state = orderInfo.State == OrderStateNew || orderInfo.PaymentCode == "offline" && orderInfo.State == OrderStatePay
		//平台收款
	case "system_receive_pay":
		state = orderInfo.State == OrderStateNew && orderInfo.PaymentCode == "online"
		//买家投诉
	case "complain":
		state = funk.ContainsInt([]int{OrderStatePay, OrderStateSend}, orderInfo.State) || orderInfo.FinnshedTime > (time.Now().Unix()-ComplainTimeLimit)
		//调整运费
	case "modify_price":
		state = orderInfo.State == OrderStateNew || orderInfo.PaymentCode == "offline" && orderInfo.State == OrderStatePay
		//state = orderInfo.ShippingFee > 0 && state
		//发货
	case "send":
		state = orderInfo.LockState == 0 && orderInfo.State == OrderStatePay
		//收货
	case "receive":
		state = orderInfo.State == OrderStateSend
		//评价
	case "evaluate":
		state = orderInfo.LockState == 0 && orderInfo.EvaluateState == 0 && orderInfo.State == OrderStateSuccess && time.Now().Unix()-orderInfo.FinnshedTime < ORDER_EVALUATE_TIME
		// 子商品是否可评价
	case "evaluate_goods":
		state = orderInfo.State == OrderStateSuccess
		//锁定
	case "lock":
		state = orderInfo.LockState == 1
		//快递跟踪
	case "deliver":
		//state = orderInfo.TrackingNo && util.InArrayInt(orderInfo.State, []int{OrderStateSend, OrderStateSuccess})
		//分享
	case "share":
		state = orderInfo.State == OrderStateSuccess
	}
	return state
}

// GetOrderFullsPage 改写getOrderList
func (f *order) GetOrderFullsPage(c *gin.Context, reqList trans.ReqPage, conds mysql.Conds, extend map[string]bool) (int, map[int]*OrderFull) {
	total, orders := Order.ListPage(c, conds, reqList)
	tmpMap := make(map[int]*OrderFull, 0)
	orderIds := make([]int, 0, total)
	uids := make([]int, 0, total)
	uidstmp := make(map[int]bool, 0)
	for _, v := range orders {
		ret := &OrderFull{Order: v}
		// 1默认2拼团商品3限时折扣商品4组合套装5赠品

		ret.StateDesc = v.GetOrderState()
		ret.StateDesc = v.GetOrderGroupState()
		ret.PaymentName = v.GetOrderPaymentName()
		tmpMap[v.Id] = ret

		orderIds = append(orderIds, v.Id)
		if _, ok := uidstmp[v.Uid]; !ok {
			uids = append(uids, v.Uid)
		}
	}

	if _, ok := extend["user"]; ok {
		users, _ := UserOpen.ListMap(c, mysql.Conds{
			"id": mysql.Cond{"in", uids},
		})
		for oid := range tmpMap {
			ret := tmpMap[oid]
			ret.UserOpen = users[ret.Order.Uid]
			tmpMap[oid] = ret
		}
	}

	if _, ok := extend["order_extend"]; ok {
		extendOrderGoods, _ := OrderExtend.ListMap(c, mysql.Conds{
			"id": mysql.Cond{"in", orderIds},
		})
		for oid := range tmpMap {
			ret := tmpMap[oid]
			ret.OrderExtend = extendOrderGoods[oid]
			tmpMap[oid] = ret
		}
	}

	if _, ok := extend["order_goods"]; ok {
		orderGoods, _ := OrderGoods.ListMap(c, mysql.Conds{
			"id": mysql.Cond{"in", orderIds},
		})
		for oid := range tmpMap {
			ret := tmpMap[oid]
			if ret.ExtendOrderGoods == nil {
				ret.ExtendOrderGoods = make([]ExtendOrderGoods, 0)
			}
			ret.ExtendOrderGoods = append(ret.ExtendOrderGoods, ExtendOrderGoods{OrderGoods: orderGoods[oid]})
			tmpMap[oid] = ret
		}
	}

	return total, tmpMap
}

// 买家订单状态操作
func (o *order) UserChangeState(c *gin.Context, stateType string, orderInfo mysql.Order, uid int, username string, remark string) bool {
	tx := model.Db.Begin()
	if stateType == "order_cancel" {
		err := o.userChangeStateOrderCancel(c, tx, orderInfo, uid, username)
		if err != nil {
			tx.Rollback()
			return false
		}
	} else if stateType == "order_receive" {
		err := o.userChangeStateOrderReceive(c, tx, orderInfo, uid, username)
		if err != nil {
			tx.Rollback()
			return false
		}
	}
	tx.Commit()
	return true
}

// 取消订单操作
func (o *order) userChangeStateOrderCancel(c *gin.Context, tx *gorm.DB, orderInfo mysql.Order, uid int, username string) (err error) {
	isAllow := o.GetOrderOperateState("user_cancel", orderInfo)
	if !isAllow {
		err = errors.New("异常")
		return
	}

	_, orderGoodsList := OrderGoods.ListPage(c, mysql.Conds{
		"order_id": orderInfo.Id,
	}, trans.ReqPage{
		Current:  0,
		PageSize: 1000,
	})

	if len(orderGoodsList) > 0 {
		for _, goods := range orderGoodsList {
			updateData := make(mysql.Ups, 0)
			updateGoodsSkuData := make(mysql.Ups, 0)
			updateData["stock"] = gorm.Expr("stock+?", goods.GoodsNum)
			updateData["sale_num"] = gorm.Expr("sale_num-?", goods.GoodsNum)
			err = Goods.UpdateX(c, tx, mysql.Conds{"id": goods.Id}, updateData)
			if err != nil {
				return
			}

			updateGoodsSkuData["stock"] = gorm.Expr("stock+?", goods.GoodsNum)
			updateGoodsSkuData["sale_num"] = gorm.Expr("sale_num-?", goods.GoodsNum)

			err = GoodsSku.UpdateX(c, tx, mysql.Conds{"id": goods.GoodsSkuId}, updateData)
			if err != nil {
				return
			}
		}
	}

	if orderInfo.PdAmount > 0 {
		err = YiPay.ChangePrePay(c, "order_cancel", uid, username, orderInfo.PdAmount, orderInfo.Sn, 0, "")
		if err != nil {
			return
		}
	}

	// 更新订单消息
	err = o.UpdateX(c, tx, mysql.Conds{
		"id": orderInfo.Id,
	}, mysql.Ups{
		"pd_amount": 0,
		"state":     common.OrderStateCancel,
	})

	// 添加订单日志
	createOrderLog := &mysql.OrderLog{
		OrderId:    orderInfo.Id,
		Msg:        "取消了订单",
		CreateTime: time.Now().Unix(),
		Role:       "买家",
		User:       "",
		OrderState: common.OrderStateCancel,
		DeleteTime: 0,
	}

	err = OrderLog.Create(c, tx, createOrderLog)
	return

}

// 收货操作

func (o *order) userChangeStateOrderReceive(c *gin.Context, tx *gorm.DB, orderInfo mysql.Order, uid int, username string) (err error) {
	isAllow := o.GetOrderOperateState("receive", orderInfo)
	if !isAllow {
		err = errors.New("异常")
		return
	}

	//更新订单状态
	err = o.UpdateX(c, tx, mysql.Conds{
		"id": orderInfo.Id,
	}, mysql.Ups{
		"finished_time": time.Now().Unix(),
		"state":         common.OrderStateSuccess,
	})
	if err != nil {
		return
	}

	// 添加日志
	createOrderLog := &mysql.OrderLog{
		OrderId:    orderInfo.Id,
		Msg:        "签收了货物",
		CreateTime: time.Now().Unix(),
		Role:       "买家",
		User:       "",
		OrderState: common.OrderStateSuccess,
	}
	err = OrderLog.Create(c, tx, createOrderLog)
	if err != nil {
		return
	}

	_, err = OrderRefund.InfoX(c, mysql.Conds{
		"order_id":    orderInfo.Id,
		"is_close":    0,
		"delete_time": 0,
	})
	//查询收货订单是否存在退款记录 存在则关闭
	if err != nil {
		//不存在不回滚
		err = nil
		return
	}

	err = OrderRefund.UpdateX(c, tx, mysql.Conds{
		"order_id": orderInfo.Id,
	}, mysql.Ups{
		"handle_time":    time.Now().Unix(),
		"handle_message": "因您确认收货，本次申请关闭",
		"is_close":       1,  //此退款关闭
		"order_lock":     1,  //锁定类型:1为不用锁定,2为需要锁定
		"handle_state":   51, //平台处理状态 默认0处理中(未处理) 10拒绝(驳回) 20同意 30成功(已完成) 50取消(用户主动撤销) 51取消(用户主动收货)
	})

	if err != nil {
		return
	}
	// 更改订单状态 解锁 子订单解锁
	err = OrderGoods.UpdateX(c, tx, mysql.Conds{
		"lock_state": 1,
		"order_id":   orderInfo.Id,
	}, mysql.Ups{
		"lock_state":          0,
		"refund_handle_state": 51,
		"refund_id":           0,
	})

	if err != nil {
		return
	}

	err = o.UpdateX(c, tx, mysql.Conds{
		"id":         orderInfo.Id,
		"lock_state": mysql.Cond{">=", 1},
	}, mysql.Ups{
		"refund_state": 0, //退款状态:0是无退款,1是部分退款,2是全部退款
		"lock_state":   0,
		"delay_time":   time.Now().Unix(),
	})
	return

}

func (o *order) Pay(c *gin.Context, paySn, paymentCode, tradeNO string, openId int) (err error) {
	tx := model.Db.Begin()

	_, err = OrderPay.InfoX(c, mysql.Conds{
		"pay_sn":    paySn,
		"pay_state": 0,
	})
	if err != nil && err == gorm.ErrRecordNotFound {
		err = errors.New("订单支付信息不存在")
		tx.Rollback()
		return
	}

	order, err := o.InfoX(c, mysql.Conds{
		"pay_sn":  paySn,
		"open_id": openId,
	})
	if err != nil && err == gorm.ErrRecordNotFound {
		err = errors.New("该订单不存在")
		tx.Rollback()
		return
	}

	err = OrderPay.UpdateX(c, tx, mysql.Conds{
		"pay_sn": paySn,
	}, mysql.Ups{
		"pay_state": 1,
	})
	if err != nil {
		err = errors.New("更新订单支付状态失败")
		tx.Rollback()
		return
	}

	// 修改订单
	orderUpdate := map[string]interface{}{
		"state":          OrderStatePay,
		"payment_time":   time.Now().Unix(),
		"payment_code":   paymentCode,
		"trade_no":       tradeNO,
		"out_request_no": fmt.Sprintf("HZ01RF00%d", order.Id),
	}

	//判断是否为拼团订单
	if order.GoodsType == 2 {
		groupState := GroupStateSuccess
		groupingOrders, err := o.List(c, mysql.Conds{"group_sign": order.GroupSign, "state": OrderStatePay, "group_state": 2})
		if err == nil && order.GroupPeopleNum == order.GroupMenNum && order.GroupMenNum == (len(groupingOrders)+1) {
			groupState = GroupStateFail
		}

		orderUpdate["group_state"] = groupState
	}

	err = o.UpdateX(c, tx, mysql.Conds{
		"pay_sn": paySn,
		"state":  OrderStateNew,
	}, orderUpdate)
	if err != nil {
		err = errors.New("更新订单状态失败")
		tx.Rollback()
		return
	}

	err = OrderLog.Create(c, tx, &mysql.OrderLog{
		OrderId:    order.Id,
		Role:       "buyer",
		Msg:        fmt.Sprintf("支付成功，支付平台交易号 : %s", tradeNO),
		OrderState: OrderStatePay,
	})
	if err != nil {
		err = errors.New("记录订单日志出现错误")
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}
