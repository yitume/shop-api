package mysql

type GoodsSpecValue struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"`     // 规格值id
	SpecId     int    `json:"spec_id" form:"spec_id" gorm:"index"` // 规格id
	Name       string `json:"name" form:"name" `                   // 规格值名称
	Sort       int    `json:"sort" form:"sort" `                   // 排序
	Color      string `json:"color" form:"color" `                 // 色彩值
	Img        string `json:"img" form:"img" `                     // 图片
	DeleteTime int64  `json:"delete_time" form:"delete_time" `     // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `             //

}

func (*GoodsSpecValue) TableName() string {
	return "goods_spec_value"
}
