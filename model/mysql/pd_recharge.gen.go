package mysql

type PdRecharge struct {
	Id           int     `json:"id" form:"id" gorm:"primary_key"`     // 自增编号
	Sn           string  `json:"sn" form:"sn" `                       // 记录唯一标示
	Uid          int     `json:"uid" form:"uid" `                     // 会员编号
	Username     string  `json:"username" form:"username" `           // 会员名称
	Amount       float64 `json:"amount" form:"amount" `               // 充值金额
	PaymentCode  string  `json:"payment_code" form:"payment_code" `   // 支付方式 alipay wxpayweb wxpayapp...
	PaymentName  string  `json:"payment_name" form:"payment_name" `   // 支付方式 支付宝 微信支付
	TradeSn      string  `json:"trade_sn" form:"trade_sn" `           // 第三方支付接口交易号
	PaymentState int     `json:"payment_state" form:"payment_state" ` // 支付状态 0未支付1支付
	PaymentTime  int64   `json:"payment_time" form:"payment_time" `   // 支付时间
	AdminName    string  `json:"admin_name" form:"admin_name" `       // 管理员名
	CreateTime   int64   `json:"create_time" form:"create_time" `     // 添加时间
	DeleteTime   int64   `json:"delete_time" form:"delete_time" `     // 软删除时间
	OpenId       int     `json:"open_id" form:"open_id" `             // 商家ID

}

func (*PdRecharge) TableName() string {
	return "pd_recharge"
}
