package common

const (
	// OrderStateCancel 取消订单
	OrderStateCancel = 0
	// OrderStateClose 关闭订单
	OrderStateClose = 0
	// OrderStateNew 未支付订单
	OrderStateNew = 10
	// OrderStatePay 已支付
	OrderStatePay = 20
	// OrderStateSend 已发货
	OrderStateSend = 30
	// OrderStateSuccess 已收货，交易成功
	OrderStateSuccess = 40
	// OrderStateUnevaluate 未评价
	OrderStateUnevaluate = 40
)

const (
	// GroupStateNew 待付款
	GroupStateNew = 0
	// GroupStatePay 正在进行中(待开团)
	GroupStatePay = 1
	// GroupStateSuccess 拼团成功
	GroupStateSuccess = 2
	// GroupStateFail 拼团失败
	GroupStateFail = 3
)

var OrderStates = map[string]int{
	"state_close":      1, // TODO
	"state_new":        OrderStateNew,
	"state_pay":        OrderStatePay,
	"state_send":       OrderStateSend,
	"state_success":    OrderStateSuccess,
	"state_cancel":     OrderStateCancel,
	"state_refund":     1,                    // TODO user表handle_state
	"state_unevaluate": OrderStateUnevaluate, // TODO user表evaluate_state
}
