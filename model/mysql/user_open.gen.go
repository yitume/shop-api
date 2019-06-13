package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type UserOpen struct {
	Uid           int                       `json:"uid" form:"uid" gorm:"primary_key"`               // ID
	Genre         int                       `json:"genre" form:"genre" `                             // 类型 1微信(app 小程序 公众号 unionid区分) 2QQ 3微博....
	WechatOpenid  string                    `json:"wechat_openid" form:"wechat_openid" gorm:"index"` // openid (如果类型是微信 则代表公众号openid)
	AppOpenid     string                    `json:"app_openid" form:"app_openid" `                   // app_openid (如果类型是微信 则代表开放平台openid)
	MiniOpenid    string                    `json:"mini_openid" form:"mini_openid" `                 // mini_openid (如果类型是微信 则代表小程序openid)
	Unionid       string                    `json:"unionid" form:"unionid" `                         // unionid
	AccessToken   string                    `json:"access_token" form:"access_token" `               // access_token
	ExpiresIn     int                       `json:"expires_in" form:"expires_in" `                   // access_token过期时间
	RefreshToken  string                    `json:"refresh_token" form:"refresh_token" `             // access_token过期可用该字段刷新用户access_token
	Scope         string                    `json:"scope" form:"scope" `                             // 应用授权作用域
	Nickname      string                    `json:"nickname" form:"nickname" `                       // 用户来源平台的昵称
	Avatar        string                    `json:"avatar" form:"avatar" `                           // 头像
	Sex           int                       `json:"sex" form:"sex" `                                 // 1男0女
	Country       string                    `json:"country" form:"country" `                         // 国家
	Province      string                    `json:"province" form:"province" `                       // 省份
	City          string                    `json:"city" form:"city" `                               // 城市
	InfoAggregate UserOpenInfoAggregateJson `json:"info_aggregate" form:"info_aggregate" `           // json 个人信息集合
	State         int                       `json:"state" form:"state" `                             // 是否绑定主帐号 默认0否 1是
	CreateTime    int64                     `json:"create_time" form:"create_time" `                 // 创建时间
	DeleteTime    int64                     `json:"delete_time" form:"delete_time" `                 // 软删除时间
	Telephone     string                    `json:"telephone" form:"telephone" `                     // 电话
	Email         string                    `json:"email" form:"email" `                             // 邮箱
	Name          string                    `json:"name" form:"name" `                               // 用户在所有平台唯一昵称
	OpenId        int                       `json:"open_id" form:"open_id" `                         // 商户ID

}

func (*UserOpen) TableName() string {
	return "user_open"
}

// 请在model/mysql/addtion.go里定义UserOpenInfoAggregateJson结构体
func (c UserOpenInfoAggregateJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *UserOpenInfoAggregateJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
