package mysql

type GoodsImage struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"`       //
	GoodsId    int    `json:"goods_id" form:"goods_id" gorm:"index"` // 商品公共内容id
	ColorId    int    `json:"color_id" form:"color_id" `             // 颜色规格值id
	Img        string `json:"img" form:"img" `                       // 商品图片
	Sort       int    `json:"sort" form:"sort" `                     // 排序
	IsDefault  int    `json:"is_default" form:"is_default" `         // 默认主题，1是，0否
	DeleteTime int64  `json:"delete_time" form:"delete_time" `       // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `               // 商户id

}

func (*GoodsImage) TableName() string {
	return "goods_image"
}
