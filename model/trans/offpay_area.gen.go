package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOffpayAreaList 你可以把ReqOffpayAreaList嵌套到需要自行修改的结构体中
type ReqOffpayAreaList struct {
	ReqPage
	mysql.OffpayArea
}

// ReqOffpayAreaCreate 你可以把ReqOffpayAreaCreate或mysql.OffpayArea嵌套到需要自行修改的结构体中
type ReqOffpayAreaCreate = mysql.OffpayArea

// ReqOffpayAreaUpdate 你可以把ReqOffpayAreaUpdate或mysql.OffpayArea嵌套到需要自行修改的结构体中
type ReqOffpayAreaUpdate = mysql.OffpayArea

// ReqOffpayAreaDelete 你可以把ReqOffpayAreaDelete嵌套到需要自行修改的结构体中
type ReqOffpayAreaDelete struct {
}
