package mysql

type Dispatch struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` //
	SnId       string `json:"sn_id" form:"sn_id" `             // 订单号
	Time       int64  `json:"time" form:"time" `               // 配送日期
	DriverId   int    `json:"driver_id" form:"driver_id" `     // 司机id
	DriverName string `json:"driver_name" form:"driver_name" ` // 司机姓名
	Number     int    `json:"number" form:"number" `           // 订单数量
	Cost       int    `json:"cost" form:"cost" `               // 费用
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商户id

}

func (*Dispatch) TableName() string {
	return "dispatch"
}
