package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderRefundReasonList 你可以把ReqOrderRefundReasonList嵌套到需要自行修改的结构体中
type ReqOrderRefundReasonList struct {
	ReqPage
	mysql.OrderRefundReason
}

// ReqOrderRefundReasonCreate 你可以把ReqOrderRefundReasonCreate或mysql.OrderRefundReason嵌套到需要自行修改的结构体中
type ReqOrderRefundReasonCreate = mysql.OrderRefundReason

// ReqOrderRefundReasonUpdate 你可以把ReqOrderRefundReasonUpdate或mysql.OrderRefundReason嵌套到需要自行修改的结构体中
type ReqOrderRefundReasonUpdate = mysql.OrderRefundReason

// ReqOrderRefundReasonDelete 你可以把ReqOrderRefundReasonDelete嵌套到需要自行修改的结构体中
type ReqOrderRefundReasonDelete struct {
	Id int `json:"id"`
}
