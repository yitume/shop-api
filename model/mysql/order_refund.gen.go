package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type OrderRefund struct {
	Id               int                           `json:"id" form:"id" gorm:"primary_key"`               // 记录ID
	OrderId          int                           `json:"order_id" form:"order_id" `                     // 订单ID
	OrderSn          string                        `json:"order_sn" form:"order_sn" `                     // 订单编号
	OrderState       int                           `json:"order_state" form:"order_state" `               // 订单状态:20:已付款;30:已发货;40:已收货;
	OrderGoodsId     int                           `json:"order_goods_id" form:"order_goods_id" `         // 子订单ID
	RefundSn         string                        `json:"refund_sn" form:"refund_sn" `                   // 申请编号
	UserId           int                           `json:"user_id" form:"user_id" `                       // 买家ID
	UserName         string                        `json:"user_name" form:"user_name" `                   // 买家会员名
	GoodsId          int                           `json:"goods_id" form:"goods_id" `                     // 商品id
	GoodsSkuId       int                           `json:"goods_sku_id" form:"goods_sku_id" `             // 商品ID,全部退款是0
	GoodsTitle       string                        `json:"goods_title" form:"goods_title" `               // 商品名称
	GoodsSpec        OrderRefundGoodsSpecJson      `json:"goods_spec" form:"goods_spec" `                 // 商品规格
	GoodsImg         string                        `json:"goods_img" form:"goods_img" `                   // 商品图片
	GoodsPrice       float64                       `json:"goods_price" form:"goods_price" `               // 商品价格 just价格
	GoodsPayPrice    float64                       `json:"goods_pay_price" form:"goods_pay_price" `       // 商品实际成交价
	GoodsNum         int                           `json:"goods_num" form:"goods_num" `                   // 商品数量
	GoodsFreightFee  float64                       `json:"goods_freight_fee" form:"goods_freight_fee" `   // 运费金额
	RefundType       int                           `json:"refund_type" form:"refund_type" `               // 申请类型:1为仅退款,2为退货退款,默认为1
	RefundAmount     float64                       `json:"refund_amount" form:"refund_amount" `           // 退款金额 小于等于子订单的金额
	OrderAmount      float64                       `json:"order_amount" form:"order_amount" `             // 总订单的金额
	OrderLock        int                           `json:"order_lock" form:"order_lock" `                 // 订单锁定类型:1为不用锁定,2为需要锁定,默认为1
	CreateTime       int64                         `json:"create_time" form:"create_time" `               // 添加时间
	UserReason       string                        `json:"user_reason" form:"user_reason" `               // 退款原因
	UserExplain      string                        `json:"user_explain" form:"user_explain" `             // 退款说明
	TrackingNo       string                        `json:"tracking_no" form:"tracking_no" `               // 退款退货物 买家发货流单号
	TrackingPhone    string                        `json:"tracking_phone" form:"tracking_phone" `         // 退款退货物 买家发货手机号
	TrackingCompany  string                        `json:"tracking_company" form:"tracking_company" `     // 退款退货 买家发货公司
	TrackingExplain  string                        `json:"tracking_explain" form:"tracking_explain" `     // 退款退货物 买家发货说明
	TrackingImages   OrderRefundTrackingImagesJson `json:"tracking_images" form:"tracking_images" `       // 退款退货物 买家发货凭证 最多6张
	TrackingTime     int64                         `json:"tracking_time" form:"tracking_time" `           // 退款退货物 买家发货时间,默认为0
	Receive          int                           `json:"receive" form:"receive" `                       // 卖家是否收到买家退货退款货物 1未收到货 2已收到货
	ReceiveTime      int64                         `json:"receive_time" form:"receive_time" `             // 卖家收货时间,默认为0
	ReceiveMessage   string                        `json:"receive_message" form:"receive_message" `       // 卖家收货备注
	PaymentCode      string                        `json:"payment_code" form:"payment_code" `             // 支付方式
	TradeNo          string                        `json:"trade_no" form:"trade_no" `                     // 支付宝交易号OR微信交易号
	OutRequestNo     string                        `json:"out_request_no" form:"out_request_no" `         // 支付宝标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	HandleState      int                           `json:"handle_state" form:"handle_state" `             // 卖家处理状态 默认0处理中(未处理) 10拒绝(驳回) 20同意 30成功(已完成) 50取消(用户主动撤销) 51取消(用户主动收货)
	HandleMessage    string                        `json:"handle_message" form:"handle_message" `         // 卖家处理信息
	HandleTime       int64                         `json:"handle_time" form:"handle_time" `               // 卖家处理时间
	BatchNo          string                        `json:"batch_no" form:"batch_no" `                     // 支付宝退款批次号 退款回调使用
	SuccessTime      int64                         `json:"success_time" form:"success_time" `             // 退款回调完成时间
	UserReceive      int                           `json:"user_receive" form:"user_receive" `             // 用户选择是否收到货物 默认0未发货(Order state=20) 1未收到货 2已收到货 卖家未发货(已付款)时无需传此参数
	UserImages       OrderRefundUserImagesJson     `json:"user_images" form:"user_images" `               // 照片凭证 json
	IsClose          int                           `json:"is_close" form:"is_close" `                     // 默认0 1已关闭(退款关闭)
	DeleteTime       int64                         `json:"delete_time" form:"delete_time" `               // 软删除时间
	HandleExpiryTime int64                         `json:"handle_expiry_time" form:"handle_expiry_time" ` // 过期时间之管理员处理
	SendExpiryTime   int64                         `json:"send_expiry_time" form:"send_expiry_time" `     // 过期时间之用户发货
	RefundNo         string                        `json:"refund_no" form:"refund_no" `                   // 退款交易号
	OpenId           int                           `json:"open_id" form:"open_id" `                       // 商家ID

}

func (*OrderRefund) TableName() string {
	return "order_refund"
}

// 请在model/mysql/addtion.go里定义OrderRefundGoodsSpecJson结构体
func (c OrderRefundGoodsSpecJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *OrderRefundGoodsSpecJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义OrderRefundTrackingImagesJson结构体
func (c OrderRefundTrackingImagesJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *OrderRefundTrackingImagesJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义OrderRefundUserImagesJson结构体
func (c OrderRefundUserImagesJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *OrderRefundUserImagesJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
