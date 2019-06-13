package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqSendAreaList 你可以把ReqSendAreaList嵌套到需要自行修改的结构体中
type ReqSendAreaList struct {
	ReqPage
	mysql.SendArea
}

// ReqSendAreaCreate 你可以把ReqSendAreaCreate或mysql.SendArea嵌套到需要自行修改的结构体中
type ReqSendAreaCreate = mysql.SendArea

// ReqSendAreaUpdate 你可以把ReqSendAreaUpdate或mysql.SendArea嵌套到需要自行修改的结构体中
type ReqSendAreaUpdate = mysql.SendArea

// ReqSendAreaDelete 你可以把ReqSendAreaDelete嵌套到需要自行修改的结构体中
type ReqSendAreaDelete struct {
	Id int `json:"id"`
}
