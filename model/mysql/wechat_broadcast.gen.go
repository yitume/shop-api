package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type WechatBroadcast struct {
	Id              int                            `json:"id" form:"id" gorm:"primary_key"`             //
	CreateTime      int64                          `json:"create_time" form:"create_time" `             //
	Condition       WechatBroadcastConditionJson   `json:"condition" form:"condition" `                 //
	SendTime        int64                          `json:"send_time" form:"send_time" `                 //
	SendContent     WechatBroadcastSendContentJson `json:"send_content" form:"send_content" `           //
	SendState       int                            `json:"send_state" form:"send_state" `               // 是否已经发生 0未发送 1成功 2失败
	SendUserCount   int                            `json:"send_user_count" form:"send_user_count" `     // 发送人数
	Openids         WechatBroadcastOpenidsJson     `json:"openids" form:"openids" `                     // 发送的openid集合
	SendContentType string                         `json:"send_content_type" form:"send_content_type" ` // `text`文本、`image`图片、`news`图文 、`voice`音频、 `video`视频
	ConditionType   int                            `json:"condition_type" form:"condition_type" `       // 1全部粉丝  2按条件筛选 3手动选择粉丝
	OpenId          int                            `json:"open_id" form:"open_id" `                     // 商家ID

}

func (*WechatBroadcast) TableName() string {
	return "wechat_broadcast"
}

// 请在model/mysql/addtion.go里定义WechatBroadcastConditionJson结构体
func (c WechatBroadcastConditionJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *WechatBroadcastConditionJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义WechatBroadcastSendContentJson结构体
func (c WechatBroadcastSendContentJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *WechatBroadcastSendContentJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义WechatBroadcastOpenidsJson结构体
func (c WechatBroadcastOpenidsJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *WechatBroadcastOpenidsJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
