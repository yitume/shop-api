package mysql

type VerifyCode struct {
	Id          int    `json:"id" form:"id" gorm:"primary_key"`   //
	Code        string `json:"code" form:"code" `                 // 验证码
	Receiver    string `json:"receiver" form:"receiver" `         // 接收者
	ChannelType string `json:"channel_type" form:"channel_type" ` // 渠道类型，sms 短信、email 邮箱
	Behavior    string `json:"behavior" form:"behavior" `         // 行为标识，如register
	Status      int    `json:"status" form:"status" `             // 状态0未验证，1已验证
	UserId      int    `json:"user_id" form:"user_id" `           // 用户id
	CreateTime  int64  `json:"create_time" form:"create_time" `   // 时间
	ExpireTime  int64  `json:"expire_time" form:"expire_time" `   // 过期时间
	Ip          string `json:"ip" form:"ip" `                     //
	SendState   int    `json:"send_state" form:"send_state" `     // 发送状态 1成功 0失败
	OpenId      int    `json:"open_id" form:"open_id" `           // 商家ID

}

func (*VerifyCode) TableName() string {
	return "verify_code"
}
