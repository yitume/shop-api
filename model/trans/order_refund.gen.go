package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderRefundList 你可以把ReqOrderRefundList嵌套到需要自行修改的结构体中
type ReqOrderRefundList struct {
	ReqPage
	mysql.OrderRefund
}

// ReqOrderRefundCreate 你可以把ReqOrderRefundCreate或mysql.OrderRefund嵌套到需要自行修改的结构体中
type ReqOrderRefundCreate = mysql.OrderRefund

// ReqOrderRefundUpdate 你可以把ReqOrderRefundUpdate或mysql.OrderRefund嵌套到需要自行修改的结构体中
type ReqOrderRefundUpdate = mysql.OrderRefund

// ReqOrderRefundDelete 你可以把ReqOrderRefundDelete嵌套到需要自行修改的结构体中
type ReqOrderRefundDelete struct {
	Id int `json:"id"`
}
