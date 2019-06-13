package mysql

type Biz struct {
	OpenId          int    `json:"open_id" form:"open_id" gorm:"primary_key"` //
	Telephone       string `json:"telephone" form:"telephone" `               //
	Nickname        string `json:"nickname" form:"nickname" `                 //
	Domain          string `json:"domain" form:"domain" `                     //
	PlatformNo      string `json:"platform_no" form:"platform_no" `           //
	PlatformSercret string `json:"platform_sercret" form:"platform_sercret" ` //
	CreatedAt       int64  `json:"created_at" form:"created_at" `             //
	UpdatedAt       int64  `json:"updated_at" form:"updated_at" `             //
	Password        string `json:"password" form:"password" `                 //
	Email           string `json:"email" form:"email" `                       //
	Avatar          string `json:"avatar" form:"avatar" `                     //
	LastLoginIp     string `json:"last_login_ip" form:"last_login_ip" `       //
	Status          int    `json:"status" form:"status" `                     //

}

func (*Biz) TableName() string {
	return "biz"
}
