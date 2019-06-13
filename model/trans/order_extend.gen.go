package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderExtendList 你可以把ReqOrderExtendList嵌套到需要自行修改的结构体中
type ReqOrderExtendList struct {
	ReqPage
	mysql.OrderExtend
}

// ReqOrderExtendCreate 你可以把ReqOrderExtendCreate或mysql.OrderExtend嵌套到需要自行修改的结构体中
type ReqOrderExtendCreate = mysql.OrderExtend

// ReqOrderExtendUpdate 你可以把ReqOrderExtendUpdate或mysql.OrderExtend嵌套到需要自行修改的结构体中
type ReqOrderExtendUpdate = mysql.OrderExtend

// ReqOrderExtendDelete 你可以把ReqOrderExtendDelete嵌套到需要自行修改的结构体中
type ReqOrderExtendDelete struct {
	Id int `json:"id"`
}
