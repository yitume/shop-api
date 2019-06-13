package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqCartList 你可以把ReqCartList嵌套到需要自行修改的结构体中
type ReqCartList struct {
	ReqPage
	mysql.Cart
}

// ReqCartCreate 你可以把ReqCartCreate或mysql.Cart嵌套到需要自行修改的结构体中
type ReqCartCreate = mysql.Cart

// ReqCartUpdate 你可以把ReqCartUpdate或mysql.Cart嵌套到需要自行修改的结构体中
type ReqCartUpdate = mysql.Cart

// ReqCartDelete 你可以把ReqCartDelete嵌套到需要自行修改的结构体中
type ReqCartDelete struct {
	Id int `json:"id"`
}
