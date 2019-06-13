package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserVisitList 你可以把ReqUserVisitList嵌套到需要自行修改的结构体中
type ReqUserVisitList struct {
	ReqPage
	mysql.UserVisit
}

// ReqUserVisitCreate 你可以把ReqUserVisitCreate或mysql.UserVisit嵌套到需要自行修改的结构体中
type ReqUserVisitCreate = mysql.UserVisit

// ReqUserVisitUpdate 你可以把ReqUserVisitUpdate或mysql.UserVisit嵌套到需要自行修改的结构体中
type ReqUserVisitUpdate = mysql.UserVisit

// ReqUserVisitDelete 你可以把ReqUserVisitDelete嵌套到需要自行修改的结构体中
type ReqUserVisitDelete struct {
}
