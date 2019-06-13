package mysql

type OrderPay struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` //
	PaySn      string `json:"pay_sn" form:"pay_sn" `           // 支付单号
	UserId     int    `json:"user_id" form:"user_id" `         // 买家ID
	PayState   int    `json:"pay_state" form:"pay_state" `     // 0默认未支付1已支付(只有第三方支付接口通知到时才会更改此状态)
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*OrderPay) TableName() string {
	return "order_pay"
}
