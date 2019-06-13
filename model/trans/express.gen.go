package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqExpressList 你可以把ReqExpressList嵌套到需要自行修改的结构体中
type ReqExpressList struct {
	ReqPage
	mysql.Express
}

// ReqExpressCreate 你可以把ReqExpressCreate或mysql.Express嵌套到需要自行修改的结构体中
type ReqExpressCreate = mysql.Express

// ReqExpressUpdate 你可以把ReqExpressUpdate或mysql.Express嵌套到需要自行修改的结构体中
type ReqExpressUpdate = mysql.Express

// ReqExpressDelete 你可以把ReqExpressDelete嵌套到需要自行修改的结构体中
type ReqExpressDelete struct {
	Id int `json:"id"`
}
