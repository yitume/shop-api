package mysql

type Address struct {
	Id            int    `json:"id" form:"id" gorm:"primary_key"`       // 地址ID
	Uid           int    `json:"uid" form:"uid" gorm:"index"`           // 会员ID
	Truename      string `json:"truename" form:"truename" `             // 会员姓名
	ProvinceId    int    `json:"province_id" form:"province_id" `       // 省份id
	CityId        int    `json:"city_id" form:"city_id" `               // 市级ID
	AreaId        int    `json:"area_id" form:"area_id" `               // 地区ID
	Address       string `json:"address" form:"address" `               // 地址
	CombineDetail string `json:"combine_detail" form:"combine_detail" ` // 地区内容组合
	TelPhone      string `json:"tel_phone" form:"tel_phone" `           // 座机电话
	MobilePhone   string `json:"mobile_phone" form:"mobile_phone" `     // 手机电话
	ZipCode       string `json:"zip_code" form:"zip_code" `             // 邮编
	IsDefault     int    `json:"is_default" form:"is_default" `         // 1默认收货地址
	Type          string `json:"type" form:"type" `                     // &#39;个人&#39;,&#39;公司&#39;,&#39;其他&#39;....
	StreetId      int    `json:"street_id" form:"street_id" `           // 街道id
	DeleteTime    int64  `json:"delete_time" form:"delete_time" `       // 软删除时间
	OpenId        int    `json:"open_id" form:"open_id" `               // 商户id

}

func (*Address) TableName() string {
	return "address"
}
