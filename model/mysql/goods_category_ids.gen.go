package mysql

type GoodsCategoryIds struct {
	Id         int   `json:"id" form:"id" gorm:"primary_key"`             //
	GoodsId    int   `json:"goods_id" form:"goods_id" gorm:"index"`       // 商品id
	CategoryId int   `json:"category_id" form:"category_id" gorm:"index"` // 商品分类id
	DeleteTime int64 `json:"delete_time" form:"delete_time" `             // 软删除时间
	OpenId     int   `json:"open_id" form:"open_id" `                     // 商户id

}

func (*GoodsCategoryIds) TableName() string {
	return "goods_category_ids"
}
