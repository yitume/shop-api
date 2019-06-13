package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderList 你可以把ReqOrderList嵌套到需要自行修改的结构体中
type ReqOrderList struct {
	ReqPage
	mysql.Order
}

// ReqOrderCreate 你可以把ReqOrderCreate或mysql.Order嵌套到需要自行修改的结构体中
type ReqOrderCreate = mysql.Order

// ReqOrderUpdate 你可以把ReqOrderUpdate或mysql.Order嵌套到需要自行修改的结构体中
type ReqOrderUpdate = mysql.Order

// ReqOrderDelete 你可以把ReqOrderDelete嵌套到需要自行修改的结构体中
type ReqOrderDelete struct {
	Id int `json:"id"`
}
