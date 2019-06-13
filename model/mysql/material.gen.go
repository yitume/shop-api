package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type Material struct {
	Id         int               `json:"id" form:"id" gorm:"primary_key"` //
	Media      MaterialMediaJson `json:"media" form:"media" `             // 媒体
	Type       string            `json:"type" form:"type" `               // news、voice、video
	CreateTime int64             `json:"create_time" form:"create_time" ` //
	OpenId     int               `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Material) TableName() string {
	return "material"
}

// 请在model/mysql/addtion.go里定义MaterialMediaJson结构体
func (c MaterialMediaJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *MaterialMediaJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
