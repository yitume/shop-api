package view

import "github.com/yitume/shop-api/model/mysql"

type FaOrder struct {
	Id                int               `json:"id" gorm:"primary_key"` // 订单索引id
	StateDesc         string            `json:"state_desc"`
	GroupStateDesc    string            `json:"group_state_desc"`
	PaymentName       string            `json:"payment_name"`
	ExtendOrderExtend mysql.OrderExtend `json:"extend_order_extend"`
}

type OrderCondition struct {
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
