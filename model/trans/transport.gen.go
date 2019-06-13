package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqTransportList 你可以把ReqTransportList嵌套到需要自行修改的结构体中
type ReqTransportList struct {
	ReqPage
	mysql.Transport
}

// ReqTransportCreate 你可以把ReqTransportCreate或mysql.Transport嵌套到需要自行修改的结构体中
type ReqTransportCreate = mysql.Transport

// ReqTransportUpdate 你可以把ReqTransportUpdate或mysql.Transport嵌套到需要自行修改的结构体中
type ReqTransportUpdate = mysql.Transport

// ReqTransportDelete 你可以把ReqTransportDelete嵌套到需要自行修改的结构体中
type ReqTransportDelete struct {
	Id int `json:"id"`
}
