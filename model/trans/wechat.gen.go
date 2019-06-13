package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqWechatList 你可以把ReqWechatList嵌套到需要自行修改的结构体中
type ReqWechatList struct {
	ReqPage
	mysql.Wechat
}

// ReqWechatCreate 你可以把ReqWechatCreate或mysql.Wechat嵌套到需要自行修改的结构体中
type ReqWechatCreate = mysql.Wechat

// ReqWechatUpdate 你可以把ReqWechatUpdate或mysql.Wechat嵌套到需要自行修改的结构体中
type ReqWechatUpdate = mysql.Wechat

// ReqWechatDelete 你可以把ReqWechatDelete嵌套到需要自行修改的结构体中
type ReqWechatDelete struct {
	Id int `json:"id"`
}
