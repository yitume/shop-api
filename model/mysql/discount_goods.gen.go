package mysql

type DiscountGoods struct {
	Id         int     `json:"id" form:"id" gorm:"primary_key"`   // id
	DiscountId int     `json:"discount_id" form:"discount_id" `   // 限时折扣id
	GoodsId    int     `json:"goods_id" form:"goods_id" `         // 商品主表id
	GoodsSkuId int     `json:"goods_sku_id" form:"goods_sku_id" ` // 商品id
	Discounts  int     `json:"discounts" form:"discounts" `       // XXX折
	Minus      float64 `json:"minus" form:"minus" `               // 立减XXX元
	Price      float64 `json:"price" form:"price" `               // 打折后XXX元
	CreateTime int64   `json:"create_time" form:"create_time" `   // 创建时间
	DeleteTime int64   `json:"delete_time" form:"delete_time" `   // 软删除时间
	OpenId     int     `json:"open_id" form:"open_id" `           // 商户id

}

func (*DiscountGoods) TableName() string {
	return "discount_goods"
}
