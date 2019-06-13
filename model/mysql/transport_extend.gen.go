package mysql

type TransportExtend struct {
	Id             int     `json:"id" form:"id" gorm:"primary_key"`         // 运费模板扩展ID
	AreaId         string  `json:"area_id" form:"area_id" `                 // 市级地区ID组成的串，以，隔开，两端也有，
	TopAreaId      string  `json:"top_area_id" form:"top_area_id" `         // 省级地区ID组成的串，以，隔开，两端也有，
	AreaName       string  `json:"area_name" form:"area_name" `             // 地区name组成的串，以，隔开
	Snum           int     `json:"snum" form:"snum" `                       // 首件数量
	Sprice         float64 `json:"sprice" form:"sprice" `                   // 首件运费
	Xnum           int     `json:"xnum" form:"xnum" `                       // 续件数量
	Xprice         float64 `json:"xprice" form:"xprice" `                   // 续件运费
	IsDefault      string  `json:"is_default" form:"is_default" `           // 是否默认运费1是2否
	TransportId    int     `json:"transport_id" form:"transport_id" `       // 运费模板ID
	TransportTitle string  `json:"transport_title" form:"transport_title" ` // 运费模板
	DeleteTime     int64   `json:"delete_time" form:"delete_time" `         // 软删除时间
	OpenId         int     `json:"open_id" form:"open_id" `                 // 商家ID

}

func (*TransportExtend) TableName() string {
	return "transport_extend"
}
