package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderRefundLogList 你可以把ReqOrderRefundLogList嵌套到需要自行修改的结构体中
type ReqOrderRefundLogList struct {
	ReqPage
	mysql.OrderRefundLog
}

// ReqOrderRefundLogCreate 你可以把ReqOrderRefundLogCreate或mysql.OrderRefundLog嵌套到需要自行修改的结构体中
type ReqOrderRefundLogCreate = mysql.OrderRefundLog

// ReqOrderRefundLogUpdate 你可以把ReqOrderRefundLogUpdate或mysql.OrderRefundLog嵌套到需要自行修改的结构体中
type ReqOrderRefundLogUpdate = mysql.OrderRefundLog

// ReqOrderRefundLogDelete 你可以把ReqOrderRefundLogDelete嵌套到需要自行修改的结构体中
type ReqOrderRefundLogDelete struct {
	Id int `json:"id"`
}
