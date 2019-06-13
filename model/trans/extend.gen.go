package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqExtendList 你可以把ReqExtendList嵌套到需要自行修改的结构体中
type ReqExtendList struct {
	ReqPage
	mysql.Extend
}

// ReqExtendCreate 你可以把ReqExtendCreate或mysql.Extend嵌套到需要自行修改的结构体中
type ReqExtendCreate = mysql.Extend

// ReqExtendUpdate 你可以把ReqExtendUpdate或mysql.Extend嵌套到需要自行修改的结构体中
type ReqExtendUpdate = mysql.Extend

// ReqExtendDelete 你可以把ReqExtendDelete嵌套到需要自行修改的结构体中
type ReqExtendDelete struct {
	Id int `json:"id"`
}
