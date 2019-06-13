package mysql

type Shipper struct {
	Id            int    `json:"id" form:"id" gorm:"primary_key"`       // 物流地址id
	Name          string `json:"name" form:"name" `                     // 发货人
	ProvinceId    int    `json:"province_id" form:"province_id" `       // 省ID
	CityId        int    `json:"city_id" form:"city_id" `               // 市ID
	AreaId        int    `json:"area_id" form:"area_id" `               // 区县ID
	StreetId      int    `json:"street_id" form:"street_id" `           // 街道ID
	CombineDetail string `json:"combine_detail" form:"combine_detail" ` // 地区信息
	Address       string `json:"address" form:"address" `               // 详细地址
	ContactNumber string `json:"contact_number" form:"contact_number" ` // 联系电话
	IsDefault     int    `json:"is_default" form:"is_default" `         // 是否为默认地址，默认 0 不是、1 是
	CreateTime    int64  `json:"create_time" form:"create_time" `       // 创建时间
	UpdateTime    int64  `json:"update_time" form:"update_time" `       // 更新时间
	DeleteTime    int64  `json:"delete_time" form:"delete_time" `       // 软删除时间
	RefundDefault int    `json:"refund_default" form:"refund_default" ` // 默认退货地址 默认0否 1是
	IsSystem      int    `json:"is_system" form:"is_system" `           // 系统级 默认0自定义 1系统
	OpenId        int    `json:"open_id" form:"open_id" `               // 商家ID

}

func (*Shipper) TableName() string {
	return "shipper"
}
