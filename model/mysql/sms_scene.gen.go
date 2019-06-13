package mysql

type SmsScene struct {
	Id                 int    `json:"id" form:"id" gorm:"primary_key"`                   //
	Name               string `json:"name" form:"name" `                                 // 名称
	Sign               string `json:"sign" form:"sign" `                                 // 唯一标识
	Signature          string `json:"signature" form:"signature" `                       // 签名
	ProviderTemplateId string `json:"provider_template_id" form:"provider_template_id" ` // 服务商提供的模板id
	ProviderType       string `json:"provider_type" form:"provider_type" `               // 提供商标识
	Body               string `json:"body" form:"body" `                                 // 内容
	IsSystem           int    `json:"is_system" form:"is_system" `                       // 系统默认
	OpenId             int    `json:"open_id" form:"open_id" `                           // 商家ID

}

func (*SmsScene) TableName() string {
	return "sms_scene"
}
