package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type Freight struct {
	Id         int              `json:"id" form:"id" gorm:"primary_key"` // 模板id
	Name       string           `json:"name" form:"name" `               // 模板名称
	PayType    int              `json:"pay_type" form:"pay_type" `       // 计算方式：1 按件数 2 按重量
	Areas      FreightAreasJson `json:"areas" form:"areas" `             //
	CreateTime int64            `json:"create_time" form:"create_time" ` // 添加时间
	UpdateTime int64            `json:"update_time" form:"update_time" ` // 更新时间
	DeleteTime int64            `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int              `json:"open_id" form:"open_id" `         // 商户id

}

func (*Freight) TableName() string {
	return "freight"
}

// 请在model/mysql/addtion.go里定义FreightAreasJson结构体
func (c FreightAreasJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *FreightAreasJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
