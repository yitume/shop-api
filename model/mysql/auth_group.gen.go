package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type AuthGroup struct {
	Id         int                  `json:"id" form:"id" gorm:"primary_key"` //
	Name       string               `json:"name" form:"name" `               //
	RuleIds    AuthGroupRuleIdsJson `json:"rule_ids" form:"rule_ids" `       //
	Status     int                  `json:"status" form:"status" `           //
	DeleteTime int64                `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int                  `json:"open_id" form:"open_id" `         // 商家ID

}

func (*AuthGroup) TableName() string {
	return "auth_group"
}

// 请在model/mysql/addtion.go里定义AuthGroupRuleIdsJson结构体
func (c AuthGroupRuleIdsJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *AuthGroupRuleIdsJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
