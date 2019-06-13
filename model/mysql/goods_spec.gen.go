package mysql

type GoodsSpec struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // 规格id
	Name       string `json:"name" form:"name" `               // 规格名称
	Sort       int    `json:"sort" form:"sort" `               // 规格排序
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*GoodsSpec) TableName() string {
	return "goods_spec"
}
