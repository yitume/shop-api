package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type SendArea struct {
	Id         int                 `json:"id" form:"id" gorm:"primary_key"` //
	Title      string              `json:"title" form:"title" `             // 配送区域模板名称
	AreaIds    SendAreaAreaIdsJson `json:"area_ids" form:"area_ids" `       // 市id集合 json
	CreateTime int64               `json:"create_time" form:"create_time" ` // 添加时间
	UpdateTime int64               `json:"update_time" form:"update_time" ` // 更新时间
	DeleteTime int64               `json:"delete_time" form:"delete_time" ` // 软删除时间
	IsSystem   int                 `json:"is_system" form:"is_system" `     // 系统级 默认0自定义 1系统
	OpenId     int                 `json:"open_id" form:"open_id" `         // 商家ID

}

func (*SendArea) TableName() string {
	return "send_area"
}

// 请在model/mysql/addtion.go里定义SendAreaAreaIdsJson结构体
func (c SendAreaAreaIdsJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *SendAreaAreaIdsJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
