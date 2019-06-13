package mysql

type Cart struct {
	Id         int   `json:"id" form:"id" gorm:"primary_key"`               // 购物车id
	Uid        int   `json:"uid" form:"uid" gorm:"index"`                   // 买家id
	GoodsSkuId int   `json:"goods_sku_id" form:"goods_sku_id" gorm:"index"` // 商品id
	GoodsNum   int   `json:"goods_num" form:"goods_num" `                   // 购买商品数量
	IsCheck    int   `json:"is_check" form:"is_check" `                     // 选中状态 默认1选中 0未选中
	CreateTime int64 `json:"create_time" form:"create_time" `               //
	DeleteTime int64 `json:"delete_time" form:"delete_time" `               // 软删除时间
	OpenId     int   `json:"open_id" form:"open_id" `                       // open_id

}

func (*Cart) TableName() string {
	return "cart"
}
