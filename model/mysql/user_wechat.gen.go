package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type UserWechat struct {
	UserId        int                     `json:"user_id" form:"user_id" gorm:"primary_key"` //
	Openid        string                  `json:"openid" form:"openid" `                     //
	Sex           int                     `json:"sex" form:"sex" `                           //
	Nickname      string                  `json:"nickname" form:"nickname" `                 //
	Language      string                  `json:"language" form:"language" `                 //
	City          string                  `json:"city" form:"city" `                         //
	Province      string                  `json:"province" form:"province" `                 //
	Country       string                  `json:"country" form:"country" `                   //
	Headimgurl    string                  `json:"headimgurl" form:"headimgurl" `             //
	Unionid       string                  `json:"unionid" form:"unionid" `                   //
	Remark        string                  `json:"remark" form:"remark" `                     //
	Groupid       int                     `json:"groupid" form:"groupid" `                   //
	TagidList     UserWechatTagidListJson `json:"tagid_list" form:"tagid_list" `             //
	Subscribe     int                     `json:"subscribe" form:"subscribe" `               //
	SubscribeTime int64                   `json:"subscribe_time" form:"subscribe_time" `     //

}

func (*UserWechat) TableName() string {
	return "user_wechat"
}

// 请在model/mysql/addtion.go里定义UserWechatTagidListJson结构体
func (c UserWechatTagidListJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *UserWechatTagidListJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
