package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqMessageTypeList 你可以把ReqMessageTypeList嵌套到需要自行修改的结构体中
type ReqMessageTypeList struct {
	ReqPage
	mysql.MessageType
}

// ReqMessageTypeCreate 你可以把ReqMessageTypeCreate或mysql.MessageType嵌套到需要自行修改的结构体中
type ReqMessageTypeCreate = mysql.MessageType

// ReqMessageTypeUpdate 你可以把ReqMessageTypeUpdate或mysql.MessageType嵌套到需要自行修改的结构体中
type ReqMessageTypeUpdate = mysql.MessageType

// ReqMessageTypeDelete 你可以把ReqMessageTypeDelete嵌套到需要自行修改的结构体中
type ReqMessageTypeDelete struct {
	Id int `json:"id"`
}
