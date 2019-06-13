package mysql

type Discount struct {
	Id          int    `json:"id" form:"id" gorm:"primary_key"`   // id
	Title       string `json:"title" form:"title" `               // 名称
	StartTime   int64  `json:"start_time" form:"start_time" `     // 开始时间
	EndTime     int64  `json:"end_time" form:"end_time" `         // 结束时间
	LimitNumber int    `json:"limit_number" form:"limit_number" ` // 限购件数 默认0不限制
	Partake     string `json:"partake" form:"partake" `           // json 优惠参与 折扣discount 满减fullcut 优惠券coupon 废弃
	CreateTime  int64  `json:"create_time" form:"create_time" `   // 创建时间
	DeleteTime  int64  `json:"delete_time" form:"delete_time" `   // 软删除时间
	OpenId      int    `json:"open_id" form:"open_id" `           // 商户id

}

func (*Discount) TableName() string {
	return "discount"
}
