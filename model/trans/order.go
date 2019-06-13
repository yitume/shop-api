package trans

type FaOrderExtend struct {
	Id int `json:"id" gorm:"primary_key"` // 订单索引id

	TrackingTime int64 `json:"tracking_time" ` // 配送时间

	TrackingNo string `json:"tracking_no" ` // 物流单号

	ShipperId int `json:"shipper_id" ` // 商家物流地址id

	ExpressId int `json:"express_id" ` // 物流公司id，默认为0 代表不需要物流

	Message string `json:"message" ` // 买家留言

	VoucherPrice int `json:"voucher_price" ` // 代金券面额

	VoucherId int `json:"voucher_id" ` // 代金券id

	VoucherCode string `json:"voucher_code" ` // 代金券编码

	Remark string `json:"remark" ` // 发货备注

	ReciverName string `json:"reciver_name" ` // 收货人姓名

	ReceiverPhone string `json:"receiver_phone" ` // 收货人电话

	ReciverInfo string `json:"reciver_info" ` // 收货人其它信息 json

	ReciverProvinceId int `json:"reciver_province_id" ` // 收货人省级ID

	ReciverCityId int `json:"reciver_city_id" ` //

	ReciverAreaId int `json:"reciver_area_id" ` //

	InvoiceInfo string `json:"invoice_info" ` // 发票信息 json

	PromotionInfo string `json:"promotion_info" ` // 促销信息备注

	EvaluateTime int64 `json:"evaluate_time" ` // 评价时间

	ServiceRemarks string `json:"service_remarks" ` // 后台客服对此订单做出的备注

	DeliverName string `json:"deliver_name" ` // 配送人名字

	DeliverPhone string `json:"deliver_phone" ` // 配送人电话

	DeliverAddress string `json:"deliver_address" ` // 配送地址

	FreightRule int `json:"freight_rule" ` // 运费规则1按商品累加运费2组合运费

	NeedExpress int `json:"need_express" ` // 是否需要物流1需要0不需要
	IsRefund    int `json:"is_refund"`

	RefundStatus int `json:"refund_status"`
}
