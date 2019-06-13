package mysql

type PdCash struct {
	Id           int     `json:"id" form:"id" gorm:"primary_key"`     // 自增编号
	Sn           int64   `json:"sn" form:"sn" `                       // 记录唯一标示
	Uid          int     `json:"uid" form:"uid" `                     // 会员编号
	Username     string  `json:"username" form:"username" `           // 会员名称
	Amount       float64 `json:"amount" form:"amount" `               // 金额
	BankName     string  `json:"bank_name" form:"bank_name" `         // 收款银行
	BankNo       string  `json:"bank_no" form:"bank_no" `             // 收款账号
	BankUser     string  `json:"bank_user" form:"bank_user" `         // 开户人姓名
	CreateTime   int64   `json:"create_time" form:"create_time" `     // 添加时间
	PaymentTime  int64   `json:"payment_time" form:"payment_time" `   // 付款时间
	PaymentState int     `json:"payment_state" form:"payment_state" ` // 提现支付状态 0默认1支付完成
	PaymentAdmin string  `json:"payment_admin" form:"payment_admin" ` // 支付管理员
	DeleteTime   int64   `json:"delete_time" form:"delete_time" `     // 软删除时间
	OpenId       int     `json:"open_id" form:"open_id" `             // 商家ID

}

func (*PdCash) TableName() string {
	return "pd_cash"
}
