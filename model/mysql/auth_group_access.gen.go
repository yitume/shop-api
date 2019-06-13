package mysql

type AuthGroupAccess struct {
	Uid     int `json:"uid" form:"uid" gorm:"primary_key"`           // 用户id
	GroupId int `json:"group_id" form:"group_id" gorm:"primary_key"` // 用户组id
	OpenId  int `json:"open_id" form:"open_id" `                     // 商家ID

}

func (*AuthGroupAccess) TableName() string {
	return "auth_group_access"
}
