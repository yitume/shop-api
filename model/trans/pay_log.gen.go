package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqPayLogList 你可以把ReqPayLogList嵌套到需要自行修改的结构体中
type ReqPayLogList struct {
	ReqPage
	mysql.PayLog
}

// ReqPayLogCreate 你可以把ReqPayLogCreate或mysql.PayLog嵌套到需要自行修改的结构体中
type ReqPayLogCreate = mysql.PayLog

// ReqPayLogUpdate 你可以把ReqPayLogUpdate或mysql.PayLog嵌套到需要自行修改的结构体中
type ReqPayLogUpdate = mysql.PayLog

// ReqPayLogDelete 你可以把ReqPayLogDelete嵌套到需要自行修改的结构体中
type ReqPayLogDelete struct {
	LogId int `json:"log_id"`
}
