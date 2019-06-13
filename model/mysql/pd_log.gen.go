package mysql

type PdLog struct {
	Id              int     `json:"id" form:"id" gorm:"primary_key"`           // 自增编号
	PdId            int     `json:"pd_id" form:"pd_id" `                       // pd_cash或pd_recharge表或order表(余额付款)id 退款表ID
	PdSn            string  `json:"pd_sn" form:"pd_sn" `                       // pd_cash或pd_recharge表交易号(sn) 或 order表(余额付款 支付单号sn) 退款表refund_sn
	Uid             int     `json:"uid" form:"uid" `                           // 会员编号
	Username        string  `json:"username" form:"username" `                 // 会员名称
	AdminName       string  `json:"admin_name" form:"admin_name" `             // 管理员名称
	Type            string  `json:"type" form:"type" `                         // order_pay下单支付预存款,order_freeze下单冻结预存款,order_cancel取消订单解冻预存款,order_comb_pay下单支付被冻结的预存款,recharge充值,cash_apply申请提现冻结预存款,cash_pay提现成功,cash_del取消提现申请，解冻预存款,refund退款
	AvailableAmount float64 `json:"available_amount" form:"available_amount" ` // 可用金额变更0表示未变更
	FreezeAmount    float64 `json:"freeze_amount" form:"freeze_amount" `       // 冻结金额变更0表示未变更
	CreateTime      int64   `json:"create_time" form:"create_time" `           // 添加时间
	Remark          string  `json:"remark" form:"remark" `                     // 描述
	State           int     `json:"state" form:"state" `                       // 0未成功 1已成功
	DeleteTime      int64   `json:"delete_time" form:"delete_time" `           // 软删除时间
	OpenId          int     `json:"open_id" form:"open_id" `                   // 商家ID

}

func (*PdLog) TableName() string {
	return "pd_log"
}
