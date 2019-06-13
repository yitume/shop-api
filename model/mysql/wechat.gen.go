package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type Wechat struct {
	Id                              int                                       `json:"id" form:"id" gorm:"primary_key"`                                                 //
	Name                            string                                    `json:"name" form:"name" `                                                               // 公众号名称，填写公众号的账号名称
	Description                     string                                    `json:"description" form:"description" `                                                 // 描述，用于说明此公众号的功能及用途
	Account                         string                                    `json:"account" form:"account" `                                                         // 公众号账号， 填写公众号的账号,一般为英文账号，如：9476400@qq.com
	Original                        string                                    `json:"original" form:"original" `                                                       // 原始ID，原始ID不能修改,请谨慎填写，如：gh_9468ce6ce474
	Level                           int                                       `json:"level" form:"level" `                                                             // 类型， 1普通订阅号、2普通服务号、3认证订阅号、4认证服务号
	AppKey                          string                                    `json:"app_key" form:"app_key" `                                                         // AppId
	AppSecret                       string                                    `json:"app_secret" form:"app_secret" `                                                   // AppSecret
	Headimg                         string                                    `json:"headimg" form:"headimg" `                                                         // 头像
	Qrcode                          string                                    `json:"qrcode" form:"qrcode" `                                                           // 二维码
	AutoReplyStatus                 int                                       `json:"auto_reply_status" form:"auto_reply_status" `                                     // 被关注自动回复状态0关闭1开启
	AutoReplySubscribeReplayContent WechatAutoReplySubscribeReplayContentJson `json:"auto_reply_subscribe_replay_content" form:"auto_reply_subscribe_replay_content" ` // 被关注自动回复内容
	MchId                           string                                    `json:"mch_id" form:"mch_id" `                                                           //
	AppId                           string                                    `json:"app_id" form:"app_id" `                                                           //
	OpenId                          int                                       `json:"open_id" form:"open_id" `                                                         // 商家ID

}

func (*Wechat) TableName() string {
	return "wechat"
}

// 请在model/mysql/addtion.go里定义WechatAutoReplySubscribeReplayContentJson结构体
func (c WechatAutoReplySubscribeReplayContentJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *WechatAutoReplySubscribeReplayContentJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
