package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type WechatUser struct {
	Id            int                     `json:"id" form:"id" gorm:"primary_key"`       //
	AppId         string                  `json:"app_id" form:"app_id" `                 //
	Subscribe     int                     `json:"subscribe" form:"subscribe" `           //
	Openid        string                  `json:"openid" form:"openid" `                 //
	Nickname      string                  `json:"nickname" form:"nickname" `             //
	Sex           int                     `json:"sex" form:"sex" `                       //
	Language      string                  `json:"language" form:"language" `             //
	City          string                  `json:"city" form:"city" `                     //
	Province      string                  `json:"province" form:"province" `             //
	Country       string                  `json:"country" form:"country" `               //
	Headimgurl    string                  `json:"headimgurl" form:"headimgurl" `         //
	SubscribeTime int64                   `json:"subscribe_time" form:"subscribe_time" ` //
	Unionid       string                  `json:"unionid" form:"unionid" `               //
	Remark        string                  `json:"remark" form:"remark" `                 //
	Groupid       string                  `json:"groupid" form:"groupid" `               //
	TagidList     WechatUserTagidListJson `json:"tagid_list" form:"tagid_list" `         //
	CreateTime    int64                   `json:"create_time" form:"create_time" `       //
	UpdateTime    int64                   `json:"update_time" form:"update_time" `       //

}

func (*WechatUser) TableName() string {
	return "wechat_user"
}

// 请在model/mysql/addtion.go里定义WechatUserTagidListJson结构体
func (c WechatUserTagidListJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *WechatUserTagidListJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
