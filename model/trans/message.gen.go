package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqMessageList 你可以把ReqMessageList嵌套到需要自行修改的结构体中
type ReqMessageList struct {
	ReqPage
	mysql.Message
}

// ReqMessageCreate 你可以把ReqMessageCreate或mysql.Message嵌套到需要自行修改的结构体中
type ReqMessageCreate = mysql.Message

// ReqMessageUpdate 你可以把ReqMessageUpdate或mysql.Message嵌套到需要自行修改的结构体中
type ReqMessageUpdate = mysql.Message

// ReqMessageDelete 你可以把ReqMessageDelete嵌套到需要自行修改的结构体中
type ReqMessageDelete struct {
	Id int `json:"id"`
}
