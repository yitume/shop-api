package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderLogList 你可以把ReqOrderLogList嵌套到需要自行修改的结构体中
type ReqOrderLogList struct {
	ReqPage
	mysql.OrderLog
}

// ReqOrderLogCreate 你可以把ReqOrderLogCreate或mysql.OrderLog嵌套到需要自行修改的结构体中
type ReqOrderLogCreate = mysql.OrderLog

// ReqOrderLogUpdate 你可以把ReqOrderLogUpdate或mysql.OrderLog嵌套到需要自行修改的结构体中
type ReqOrderLogUpdate = mysql.OrderLog

// ReqOrderLogDelete 你可以把ReqOrderLogDelete嵌套到需要自行修改的结构体中
type ReqOrderLogDelete struct {
	Id int `json:"id"`
}
