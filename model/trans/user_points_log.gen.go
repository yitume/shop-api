package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserPointsLogList 你可以把ReqUserPointsLogList嵌套到需要自行修改的结构体中
type ReqUserPointsLogList struct {
	ReqPage
	mysql.UserPointsLog
}

// ReqUserPointsLogCreate 你可以把ReqUserPointsLogCreate或mysql.UserPointsLog嵌套到需要自行修改的结构体中
type ReqUserPointsLogCreate = mysql.UserPointsLog

// ReqUserPointsLogUpdate 你可以把ReqUserPointsLogUpdate或mysql.UserPointsLog嵌套到需要自行修改的结构体中
type ReqUserPointsLogUpdate = mysql.UserPointsLog

// ReqUserPointsLogDelete 你可以把ReqUserPointsLogDelete嵌套到需要自行修改的结构体中
type ReqUserPointsLogDelete struct {
	Id int `json:"id"`
}
