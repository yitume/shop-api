package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type WechatAutoReply struct {
	Id           int                             `json:"id" form:"id" gorm:"primary_key"`     //
	RuleName     string                          `json:"rule_name" form:"rule_name" `         // 规则名称
	CreateTime   int64                           `json:"create_time" form:"create_time" `     //
	ReplyMode    string                          `json:"reply_mode" form:"reply_mode" `       //
	ReplyContent WechatAutoReplyReplyContentJson `json:"reply_content" form:"reply_content" ` //
	Keys         WechatAutoReplyKeysJson         `json:"keys" form:"keys" `                   //
	OpenId       int                             `json:"open_id" form:"open_id" `             // 商家ID

}

func (*WechatAutoReply) TableName() string {
	return "wechat_auto_reply"
}

// 请在model/mysql/addtion.go里定义WechatAutoReplyReplyContentJson结构体
func (c WechatAutoReplyReplyContentJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *WechatAutoReplyReplyContentJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义WechatAutoReplyKeysJson结构体
func (c WechatAutoReplyKeysJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *WechatAutoReplyKeysJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
