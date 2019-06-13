package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type OrderGoods struct {
	Id                int                     `json:"id" form:"id" gorm:"primary_key"`                 // 订单商品表索引id
	UserId            int                     `json:"user_id" form:"user_id" `                         // 买家ID
	OrderId           int                     `json:"order_id" form:"order_id" gorm:"index"`           // 订单id
	GoodsId           int                     `json:"goods_id" form:"goods_id" `                       // 商品主表id
	GoodsSkuId        int                     `json:"goods_sku_id" form:"goods_sku_id" `               // 商品id
	GoodsTitle        string                  `json:"goods_title" form:"goods_title" `                 // 商品名称
	GoodsPrice        float64                 `json:"goods_price" form:"goods_price" `                 // 商品价格 just价格
	GoodsPayPrice     float64                 `json:"goods_pay_price" form:"goods_pay_price" `         // 商品实际支付费用(拼团商品适用)
	GoodsNum          int                     `json:"goods_num" form:"goods_num" `                     // 商品数量
	GoodsImg          string                  `json:"goods_img" form:"goods_img" `                     // 商品图片
	GoodsSpec         OrderGoodsGoodsSpecJson `json:"goods_spec" form:"goods_spec" `                   // 商品规格
	GoodsType         int                     `json:"goods_type" form:"goods_type" `                   // 1默认2团购商品3限时折扣商品4组合套装5赠品
	GoodsFreightWay   string                  `json:"goods_freight_way" form:"goods_freight_way" `     // 商品运费方式
	GoodsFreightFee   float64                 `json:"goods_freight_fee" form:"goods_freight_fee" `     // 商品的运费
	EvaluateState     int                     `json:"evaluate_state" form:"evaluate_state" `           // 评价状态 0未评价，1已评价，2已追评
	EvaluateTime      int64                   `json:"evaluate_time" form:"evaluate_time" `             // 评价时间
	CreateTime        int64                   `json:"create_time" form:"create_time" `                 // 创建时间
	CouponId          int                     `json:"coupon_id" form:"coupon_id" `                     // 线上卡券 大于0线上卡券 微信卡券表表ID 一个规格的商品对应一张微信卡券
	CouponCardId      string                  `json:"coupon_card_id" form:"coupon_card_id" `           // 线上卡券 微信卡券表微信卡券ID
	LockState         int                     `json:"lock_state" form:"lock_state" `                   // 锁定（1退款中）
	RefundHandleState int                     `json:"refund_handle_state" form:"refund_handle_state" ` // 退款平台处理状态 默认0处理中(未处理) 10拒绝(驳回) 20同意 30成功(已完成) 50取消(用户主动撤销) 51取消(用户主动收货)
	RefundId          int                     `json:"refund_id" form:"refund_id" `                     // 退款ID
	DeleteTime        int64                   `json:"delete_time" form:"delete_time" `                 // 软删除时间
	GroupPrice        float64                 `json:"group_price" form:"group_price" `                 // 拼团价格 just价格
	CaptainPrice      float64                 `json:"captain_price" form:"captain_price" `             // 团长价格 just价格
	GoodsRevisePrice  float64                 `json:"goods_revise_price" form:"goods_revise_price" `   // 修改过的商品实际支付费用 大于0起作用
	OpenId            int                     `json:"open_id" form:"open_id" `                         // 商户ID

}

func (*OrderGoods) TableName() string {
	return "order_goods"
}

// 请在model/mysql/addtion.go里定义OrderGoodsGoodsSpecJson结构体
func (c OrderGoodsGoodsSpecJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *OrderGoodsGoodsSpecJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
