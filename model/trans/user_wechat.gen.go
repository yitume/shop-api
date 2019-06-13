package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserWechatList 你可以把ReqUserWechatList嵌套到需要自行修改的结构体中
type ReqUserWechatList struct {
	ReqPage
	mysql.UserWechat
}

// ReqUserWechatCreate 你可以把ReqUserWechatCreate或mysql.UserWechat嵌套到需要自行修改的结构体中
type ReqUserWechatCreate = mysql.UserWechat

// ReqUserWechatUpdate 你可以把ReqUserWechatUpdate或mysql.UserWechat嵌套到需要自行修改的结构体中
type ReqUserWechatUpdate = mysql.UserWechat

// ReqUserWechatDelete 你可以把ReqUserWechatDelete嵌套到需要自行修改的结构体中
type ReqUserWechatDelete struct {
	UserId int `json:"user_id"`
}
