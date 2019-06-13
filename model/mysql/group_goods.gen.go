package mysql

type GroupGoods struct {
	Id           int     `json:"id" form:"id" gorm:"primary_key"`     // id
	GroupId      int     `json:"group_id" form:"group_id" `           // 拼团id
	GoodsId      int     `json:"goods_id" form:"goods_id" `           // 商品主表id
	GoodsSkuId   int     `json:"goods_sku_id" form:"goods_sku_id" `   // 商品id
	GroupPrice   float64 `json:"group_price" form:"group_price" `     // 拼团价格
	CaptainPrice float64 `json:"captain_price" form:"captain_price" ` // 团长价格
	CreateTime   int64   `json:"create_time" form:"create_time" `     // 创建时间
	UpdateTime   int64   `json:"update_time" form:"update_time" `     // 更新时间
	DeleteTime   int64   `json:"delete_time" form:"delete_time" `     // 软删除时间
	OpenId       int     `json:"open_id" form:"open_id" `             //

}

func (*GroupGoods) TableName() string {
	return "group_goods"
}
