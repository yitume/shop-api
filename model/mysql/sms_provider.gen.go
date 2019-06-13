package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type SmsProvider struct {
	Type   string                `json:"type" form:"type" gorm:"primary_key"` // 支付代码名称
	Name   string                `json:"name" form:"name" `                   // 支付名称
	Config SmsProviderConfigJson `json:"config" form:"config" `               //
	Status int                   `json:"status" form:"status" `               // 接口状态0禁用1启用
	OpenId int                   `json:"open_id" form:"open_id" `             // 商家ID

}

func (*SmsProvider) TableName() string {
	return "sms_provider"
}

// 请在model/mysql/addtion.go里定义SmsProviderConfigJson结构体
func (c SmsProviderConfigJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *SmsProviderConfigJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
