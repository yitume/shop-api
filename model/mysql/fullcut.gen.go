package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type Fullcut struct {
	Id         int                  `json:"id" form:"id" gorm:"primary_key"` // id
	Title      string               `json:"title" form:"title" `             // 名称
	StartTime  int64                `json:"start_time" form:"start_time" `   // 开始时间
	EndTime    int64                `json:"end_time" form:"end_time" `       // 结束时间
	Partake    string               `json:"partake" form:"partake" `         // json 优惠参与 折扣discount 满减fullcut 优惠券coupon 废弃
	CreateTime int64                `json:"create_time" form:"create_time" ` // 创建时间
	DeleteTime int64                `json:"delete_time" form:"delete_time" ` // 软删除时间
	Hierarchy  FullcutHierarchyJson `json:"hierarchy" form:"hierarchy" `     // 层级 至多5个,每个(包涵fll_price满XXX元,minus减XXX元,discountsXXX折,type满减类型 默认0减XXX元  1打XXX折)
	Level      int                  `json:"level" form:"level" `             // 级别 默认0全店 1商品级
	OpenId     int                  `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Fullcut) TableName() string {
	return "fullcut"
}

// 请在model/mysql/addtion.go里定义FullcutHierarchyJson结构体
func (c FullcutHierarchyJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *FullcutHierarchyJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
