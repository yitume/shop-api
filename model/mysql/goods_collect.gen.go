package mysql

type GoodsCollect struct {
	Id         int   `json:"id" form:"id" gorm:"primary_key"`   // 主键
	Uid        int   `json:"uid" form:"uid" `                   // 使用UID
	GoodsSkuId int   `json:"goods_sku_id" form:"goods_sku_id" ` // 商品skuID
	GoodsId    int   `json:"goods_id" form:"goods_id" `         // 商品id
	CreateTime int64 `json:"create_time" form:"create_time" `   // 收藏时间
	DeleteTime int64 `json:"delete_time" form:"delete_time" `   // 软删除时间
	OpenId     int   `json:"open_id" form:"open_id" `           // 商户id

}

func (*GoodsCollect) TableName() string {
	return "goods_collect"
}
