package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqQueueList 你可以把ReqQueueList嵌套到需要自行修改的结构体中
type ReqQueueList struct {
	ReqPage
	mysql.Queue
}

// ReqQueueCreate 你可以把ReqQueueCreate或mysql.Queue嵌套到需要自行修改的结构体中
type ReqQueueCreate = mysql.Queue

// ReqQueueUpdate 你可以把ReqQueueUpdate或mysql.Queue嵌套到需要自行修改的结构体中
type ReqQueueUpdate = mysql.Queue

// ReqQueueDelete 你可以把ReqQueueDelete嵌套到需要自行修改的结构体中
type ReqQueueDelete struct {
	Id int `json:"id"`
}
