package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type Setting struct {
	Key        string            `json:"key" form:"key" gorm:"primary_key"` // 键值
	Name       string            `json:"name" form:"name" `                 // 名称
	Config     SettingConfigJson `json:"config" form:"config" `             // 配置
	Status     int               `json:"status" form:"status" `             // 接口状态0禁用1启用
	Remark     string            `json:"remark" form:"remark" `             // 备注
	DeleteTime int64             `json:"delete_time" form:"delete_time" `   // 软删除时间
	CreateTime int64             `json:"create_time" form:"create_time" `   //
	UpdateTime int64             `json:"update_time" form:"update_time" `   //
	OpenId     int               `json:"open_id" form:"open_id" `           // 商家ID

}

func (*Setting) TableName() string {
	return "setting"
}

// 请在model/mysql/addtion.go里定义SettingConfigJson结构体
func (c SettingConfigJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *SettingConfigJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
