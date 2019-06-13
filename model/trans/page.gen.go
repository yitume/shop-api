package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqPageList 你可以把ReqPageList嵌套到需要自行修改的结构体中
type ReqPageList struct {
	ReqPage
	mysql.Page
}

// ReqPageCreate 你可以把ReqPageCreate或mysql.Page嵌套到需要自行修改的结构体中
type ReqPageCreate = mysql.Page

// ReqPageUpdate 你可以把ReqPageUpdate或mysql.Page嵌套到需要自行修改的结构体中
type ReqPageUpdate = mysql.Page

// ReqPageDelete 你可以把ReqPageDelete嵌套到需要自行修改的结构体中
type ReqPageDelete struct {
	Id int `json:"id"`
}
