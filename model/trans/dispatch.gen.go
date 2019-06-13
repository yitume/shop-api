package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqDispatchList 你可以把ReqDispatchList嵌套到需要自行修改的结构体中
type ReqDispatchList struct {
	ReqPage
	mysql.Dispatch
}

// ReqDispatchCreate 你可以把ReqDispatchCreate或mysql.Dispatch嵌套到需要自行修改的结构体中
type ReqDispatchCreate = mysql.Dispatch

// ReqDispatchUpdate 你可以把ReqDispatchUpdate或mysql.Dispatch嵌套到需要自行修改的结构体中
type ReqDispatchUpdate = mysql.Dispatch

// ReqDispatchDelete 你可以把ReqDispatchDelete嵌套到需要自行修改的结构体中
type ReqDispatchDelete struct {
	Id int `json:"id"`
}
