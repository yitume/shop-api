package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqWechatBroadcastList 你可以把ReqWechatBroadcastList嵌套到需要自行修改的结构体中
type ReqWechatBroadcastList struct {
	ReqPage
	mysql.WechatBroadcast
}

// ReqWechatBroadcastCreate 你可以把ReqWechatBroadcastCreate或mysql.WechatBroadcast嵌套到需要自行修改的结构体中
type ReqWechatBroadcastCreate = mysql.WechatBroadcast

// ReqWechatBroadcastUpdate 你可以把ReqWechatBroadcastUpdate或mysql.WechatBroadcast嵌套到需要自行修改的结构体中
type ReqWechatBroadcastUpdate = mysql.WechatBroadcast

// ReqWechatBroadcastDelete 你可以把ReqWechatBroadcastDelete嵌套到需要自行修改的结构体中
type ReqWechatBroadcastDelete struct {
	Id int `json:"id"`
}
