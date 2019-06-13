package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqSystemList 你可以把ReqSystemList嵌套到需要自行修改的结构体中
type ReqSystemList struct {
	ReqPage
	mysql.System
}

// ReqSystemCreate 你可以把ReqSystemCreate或mysql.System嵌套到需要自行修改的结构体中
type ReqSystemCreate = mysql.System

// ReqSystemUpdate 你可以把ReqSystemUpdate或mysql.System嵌套到需要自行修改的结构体中
type ReqSystemUpdate = mysql.System

// ReqSystemDelete 你可以把ReqSystemDelete嵌套到需要自行修改的结构体中
type ReqSystemDelete struct {
	Name string `json:"name"`
}
