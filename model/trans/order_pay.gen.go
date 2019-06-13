package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderPayList 你可以把ReqOrderPayList嵌套到需要自行修改的结构体中
type ReqOrderPayList struct {
	ReqPage
	mysql.OrderPay
}

// ReqOrderPayCreate 你可以把ReqOrderPayCreate或mysql.OrderPay嵌套到需要自行修改的结构体中
type ReqOrderPayCreate = mysql.OrderPay

// ReqOrderPayUpdate 你可以把ReqOrderPayUpdate或mysql.OrderPay嵌套到需要自行修改的结构体中
type ReqOrderPayUpdate = mysql.OrderPay

// ReqOrderPayDelete 你可以把ReqOrderPayDelete嵌套到需要自行修改的结构体中
type ReqOrderPayDelete struct {
	Id int `json:"id"`
}
