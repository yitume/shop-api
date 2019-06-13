package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqFreightList 你可以把ReqFreightList嵌套到需要自行修改的结构体中
type ReqFreightList struct {
	ReqPage
	mysql.Freight
}

// ReqFreightCreate 你可以把ReqFreightCreate或mysql.Freight嵌套到需要自行修改的结构体中
type ReqFreightCreate = mysql.Freight

// ReqFreightUpdate 你可以把ReqFreightUpdate或mysql.Freight嵌套到需要自行修改的结构体中
type ReqFreightUpdate = mysql.Freight

// ReqFreightDelete 你可以把ReqFreightDelete嵌套到需要自行修改的结构体中
type ReqFreightDelete struct {
	Id int `json:"id"`
}
