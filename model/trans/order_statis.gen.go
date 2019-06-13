package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderStatisList 你可以把ReqOrderStatisList嵌套到需要自行修改的结构体中
type ReqOrderStatisList struct {
	ReqPage
	mysql.OrderStatis
}

// ReqOrderStatisCreate 你可以把ReqOrderStatisCreate或mysql.OrderStatis嵌套到需要自行修改的结构体中
type ReqOrderStatisCreate = mysql.OrderStatis

// ReqOrderStatisUpdate 你可以把ReqOrderStatisUpdate或mysql.OrderStatis嵌套到需要自行修改的结构体中
type ReqOrderStatisUpdate = mysql.OrderStatis

// ReqOrderStatisDelete 你可以把ReqOrderStatisDelete嵌套到需要自行修改的结构体中
type ReqOrderStatisDelete struct {
	Id int64 `json:"id"`
}
