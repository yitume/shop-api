package mysql

type System struct {
	Name   string `json:"name" form:"name" gorm:"primary_key"` // fashop的当前版本
	Value  string `json:"value" form:"value" `                 //
	OpenId int    `json:"open_id" form:"open_id" `             // 商家ID

}

func (*System) TableName() string {
	return "system"
}
