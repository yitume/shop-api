package mysql

type AccessToken struct {
	Jti        int    `json:"jti" form:"jti" gorm:"primary_key"` //
	Sub        int    `json:"sub" form:"sub" `                   //
	IaTime     int64  `json:"ia_time" form:"ia_time" `           //
	ExpTime    int64  `json:"exp_time" form:"exp_time" `         // 过期时间
	Ip         string `json:"ip" form:"ip" `                     //
	CreateTime int64  `json:"create_time" form:"create_time" `   //
	IsLogout   int    `json:"is_logout" form:"is_logout" `       // 是否主动退出
	IsInvalid  int    `json:"is_invalid" form:"is_invalid" `     // 作废，当修改密码后会作废 1作废 0没有作废
	LogoutTime int64  `json:"logout_time" form:"logout_time" `   //

}

func (*AccessToken) TableName() string {
	return "access_token"
}
