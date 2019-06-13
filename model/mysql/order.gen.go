package mysql

type Order struct {
	Id                 int     `json:"id" form:"id" gorm:"primary_key"`                   // 订单索引id
	Sn                 string  `json:"sn" form:"sn" `                                     // 订单编号
	PaySn              string  `json:"pay_sn" form:"pay_sn" `                             // 支付单号
	Uid                int     `json:"uid" form:"uid" gorm:"index"`                       // 买家id
	UserName           string  `json:"user_name" form:"user_name" `                       // 买家姓名
	UserPhone          string  `json:"user_phone" form:"user_phone" `                     // 买家手机号码
	UserEmail          string  `json:"user_email" form:"user_email" `                     // 买家电子邮箱
	CreateTime         int64   `json:"create_time" form:"create_time" `                   // 订单生成时间
	PaymentCode        string  `json:"payment_code" form:"payment_code" `                 // 支付方式名称代码
	PayName            string  `json:"pay_name" form:"pay_name" `                         // 支付方式
	PaymentTime        int64   `json:"payment_time" form:"payment_time" `                 // 支付成功的(付款的)时间
	FinnshedTime       int64   `json:"finnshed_time" form:"finnshed_time" `               // 订单完成时间
	GoodsAmount        float64 `json:"goods_amount" form:"goods_amount" `                 // 商品总价格
	GoodsNum           int     `json:"goods_num" form:"goods_num" `                       // 商品个数
	Amount             float64 `json:"amount" form:"amount" `                             // 订单总价格
	PdAmount           float64 `json:"pd_amount" form:"pd_amount" `                       // 预存款支付金额
	FreightFee         float64 `json:"freight_fee" form:"freight_fee" `                   // 实际支付的运费
	FreightUnifiedFee  float64 `json:"freight_unified_fee" form:"freight_unified_fee" `   // 运费统一运费
	FreightTemplateFee float64 `json:"freight_template_fee" form:"freight_template_fee" ` // 运费模板运费
	State              int     `json:"state" form:"state" `                               // 订单状态：0(已取消)10(默认):未付款;20:已付款;30:已发货;40:已收货;
	RefundAmount       float64 `json:"refund_amount" form:"refund_amount" `               // 退款金额
	RefundState        int     `json:"refund_state" form:"refund_state" `                 // 退款状态:0是无退款,1是部分退款,2是全部退款
	LockState          int     `json:"lock_state" form:"lock_state" `                     // 锁定状态:0是正常,大于0是锁定,默认是0
	DelayTime          int64   `json:"delay_time" form:"delay_time" `                     // 延迟时间,默认为0
	Points             int     `json:"points" form:"points" `                             // 积分
	EvaluateState      int     `json:"evaluate_state" form:"evaluate_state" `             // 评价状态 0未评价，1已评价
	IsPrint            int     `json:"is_print" form:"is_print" `                         // 是否打印 1打印 0未打印
	TradeNo            string  `json:"trade_no" form:"trade_no" `                         // 支付宝交易号OR微信交易号
	OutRequestNo       string  `json:"out_request_no" form:"out_request_no" `             // 支付宝标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	WechatOpenid       string  `json:"wechat_openid" form:"wechat_openid" `               // 用户微信openid
	From               int     `json:"from" form:"from" `                                 // 1WEB2mobile
	DeleteTime         int64   `json:"delete_time" form:"delete_time" `                   // 软删除时间
	AllAgreeRefound    int     `json:"all_agree_refound" form:"all_agree_refound" `       // 默认为0，1是订单的全部商品都退了
	PayableTime        int64   `json:"payable_time" form:"payable_time" `                 // 订单可支付时间 下单时间+24小时 时间戳
	GoodsType          int     `json:"goods_type" form:"goods_type" `                     // 1默认2拼团商品3限时折扣商品4组合套装5赠品
	GroupId            int     `json:"group_id" form:"group_id" `                         // 拼团id
	GroupIdentity      int     `json:"group_identity" form:"group_identity" `             // 默认1 团长 2 团员
	GroupSign          string  `json:"group_sign" form:"group_sign" `                     // 一个团拥有相同的值（团购编号）
	GroupPeopleNum     int     `json:"group_people_num" form:"group_people_num" `         // 团共需人数 几人团
	GroupMenNum        int     `json:"group_men_num" form:"group_men_num" `               // 团现有人数
	GroupState         int     `json:"group_state" form:"group_state" `                   // 拼团状态 0待付款 1正在进行中(待开团) 2拼团成功 3拼团失败
	GroupStartTime     int64   `json:"group_start_time" form:"group_start_time" `         // 团购开始时间
	GroupEndTime       int64   `json:"group_end_time" form:"group_end_time" `             // 团购结束时间
	GoodsGroupAmount   float64 `json:"goods_group_amount" form:"goods_group_amount" `     // 商品拼团价总价格
	GroupFailTime      int64   `json:"group_fail_time" form:"group_fail_time" `           // 拼团失败时间
	ReviseAmount       float64 `json:"revise_amount" form:"revise_amount" `               // 修改过的订单总价格 大于0起作用
	ReviseFreightFee   float64 `json:"revise_freight_fee" form:"revise_freight_fee" `     // 修改过的实际支付的运费
	OpenId             int     `json:"open_id" form:"open_id" `                           //

}

func (*Order) TableName() string {
	return "order"
}
