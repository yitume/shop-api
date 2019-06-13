package mysql

type Extend struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // ID
	Express    string `json:"express" form:"express" `         // 快递公司ID的组合
	CreateTime int64  `json:"create_time" form:"create_time" ` // 时间
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Extend) TableName() string {
	return "extend"
}
