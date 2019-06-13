package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqVisitList 你可以把ReqVisitList嵌套到需要自行修改的结构体中
type ReqVisitList struct {
	ReqPage
	mysql.Visit
}

// ReqVisitCreate 你可以把ReqVisitCreate或mysql.Visit嵌套到需要自行修改的结构体中
type ReqVisitCreate = mysql.Visit

// ReqVisitUpdate 你可以把ReqVisitUpdate或mysql.Visit嵌套到需要自行修改的结构体中
type ReqVisitUpdate = mysql.Visit

// ReqVisitDelete 你可以把ReqVisitDelete嵌套到需要自行修改的结构体中
type ReqVisitDelete struct {
	Id int `json:"id"`
}
