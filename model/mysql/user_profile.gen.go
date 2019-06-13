package mysql

type UserProfile struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // ID
	Uid        int    `json:"uid" form:"uid" gorm:"index"`     // 用户id
	Name       string `json:"name" form:"name" `               // 客户姓名
	Nickname   string `json:"nickname" form:"nickname" `       // 昵称
	Avatar     string `json:"avatar" form:"avatar" `           // 头像
	Sex        int    `json:"sex" form:"sex" `                 // 1男0女
	Birthday   int    `json:"birthday" form:"birthday" `       // 生日
	Qq         string `json:"qq" form:"qq" `                   // QQ
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*UserProfile) TableName() string {
	return "user_profile"
}
