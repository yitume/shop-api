package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqMessageStateList 你可以把ReqMessageStateList嵌套到需要自行修改的结构体中
type ReqMessageStateList struct {
	ReqPage
	mysql.MessageState
}

// ReqMessageStateCreate 你可以把ReqMessageStateCreate或mysql.MessageState嵌套到需要自行修改的结构体中
type ReqMessageStateCreate = mysql.MessageState

// ReqMessageStateUpdate 你可以把ReqMessageStateUpdate或mysql.MessageState嵌套到需要自行修改的结构体中
type ReqMessageStateUpdate = mysql.MessageState

// ReqMessageStateDelete 你可以把ReqMessageStateDelete嵌套到需要自行修改的结构体中
type ReqMessageStateDelete struct {
	Id int `json:"id"`
}
