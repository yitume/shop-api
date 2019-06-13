package mysql

type Sms struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` //
	Phone      string `json:"phone" form:"phone" `             // 手机号
	CreateTime int64  `json:"create_time" form:"create_time" ` // 时间
	Status     int    `json:"status" form:"status" `           // 状态0未验证，1已验证
	Model      string `json:"model" form:"model" `             // 标识
	Text       string `json:"text" form:"text" `               // 内容
	UserId     int    `json:"user_id" form:"user_id" `         // 用户id
	UserPhone  string `json:"user_phone" form:"user_phone" `   // 用户电话
	Code       string `json:"code" form:"code" `               // 验证码
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Sms) TableName() string {
	return "sms"
}
