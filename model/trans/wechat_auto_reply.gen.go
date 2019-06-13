package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqWechatAutoReplyList 你可以把ReqWechatAutoReplyList嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyList struct {
	ReqPage
	mysql.WechatAutoReply
}

// ReqWechatAutoReplyCreate 你可以把ReqWechatAutoReplyCreate或mysql.WechatAutoReply嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyCreate = mysql.WechatAutoReply

// ReqWechatAutoReplyUpdate 你可以把ReqWechatAutoReplyUpdate或mysql.WechatAutoReply嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyUpdate = mysql.WechatAutoReply

// ReqWechatAutoReplyDelete 你可以把ReqWechatAutoReplyDelete嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyDelete struct {
	Id int `json:"id"`
}
