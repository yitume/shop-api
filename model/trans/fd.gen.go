package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqFdList 你可以把ReqFdList嵌套到需要自行修改的结构体中
type ReqFdList struct {
	ReqPage
	mysql.Fd
}

// ReqFdCreate 你可以把ReqFdCreate或mysql.Fd嵌套到需要自行修改的结构体中
type ReqFdCreate = mysql.Fd

// ReqFdUpdate 你可以把ReqFdUpdate或mysql.Fd嵌套到需要自行修改的结构体中
type ReqFdUpdate = mysql.Fd

// ReqFdDelete 你可以把ReqFdDelete嵌套到需要自行修改的结构体中
type ReqFdDelete struct {
	Id int `json:"id"`
}
