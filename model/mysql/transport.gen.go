package mysql

type Transport struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // 运费模板ID
	Title      string `json:"title" form:"title" `             // 运费模板名称
	SendTplId  int    `json:"send_tpl_id" form:"send_tpl_id" ` // 发货地区模板ID
	UpdateTime int64  `json:"update_time" form:"update_time" ` // 最后更新时间
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Transport) TableName() string {
	return "transport"
}
