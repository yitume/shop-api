package mysql

type PayLog struct {
	LogId       int     `json:"log_id" form:"log_id" gorm:"primary_key"` // 支付记录自增id
	OrderId     int     `json:"order_id" form:"order_id" `               // 对应的交交易记录的id,取值表order_info
	OrderAmount float64 `json:"order_amount" form:"order_amount" `       // 支付金额
	OrderType   int     `json:"order_type" form:"order_type" `           // 支付类型,0订单支付,1会员预付款支付
	IsPaid      int     `json:"is_paid" form:"is_paid" `                 // 是否已支付,0否;1是
	DeleteTime  int64   `json:"delete_time" form:"delete_time" `         // 软删除时间
	OpenId      int     `json:"open_id" form:"open_id" `                 // 商家ID

}

func (*PayLog) TableName() string {
	return "pay_log"
}
