package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqPdLogList 你可以把ReqPdLogList嵌套到需要自行修改的结构体中
type ReqPdLogList struct {
	ReqPage
	mysql.PdLog
}

// ReqPdLogCreate 你可以把ReqPdLogCreate或mysql.PdLog嵌套到需要自行修改的结构体中
type ReqPdLogCreate = mysql.PdLog

// ReqPdLogUpdate 你可以把ReqPdLogUpdate或mysql.PdLog嵌套到需要自行修改的结构体中
type ReqPdLogUpdate = mysql.PdLog

// ReqPdLogDelete 你可以把ReqPdLogDelete嵌套到需要自行修改的结构体中
type ReqPdLogDelete struct {
	Id int `json:"id"`
}
