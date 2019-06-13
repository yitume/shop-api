package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqWechatUserList 你可以把ReqWechatUserList嵌套到需要自行修改的结构体中
type ReqWechatUserList struct {
	ReqPage
	mysql.WechatUser
}

// ReqWechatUserCreate 你可以把ReqWechatUserCreate或mysql.WechatUser嵌套到需要自行修改的结构体中
type ReqWechatUserCreate = mysql.WechatUser

// ReqWechatUserUpdate 你可以把ReqWechatUserUpdate或mysql.WechatUser嵌套到需要自行修改的结构体中
type ReqWechatUserUpdate = mysql.WechatUser

// ReqWechatUserDelete 你可以把ReqWechatUserDelete嵌套到需要自行修改的结构体中
type ReqWechatUserDelete struct {
	Id int `json:"id"`
}
