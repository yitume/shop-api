package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqShipperList 你可以把ReqShipperList嵌套到需要自行修改的结构体中
type ReqShipperList struct {
	ReqPage
	mysql.Shipper
}

// ReqShipperCreate 你可以把ReqShipperCreate或mysql.Shipper嵌套到需要自行修改的结构体中
type ReqShipperCreate = mysql.Shipper

// ReqShipperUpdate 你可以把ReqShipperUpdate或mysql.Shipper嵌套到需要自行修改的结构体中
type ReqShipperUpdate = mysql.Shipper

// ReqShipperDelete 你可以把ReqShipperDelete嵌套到需要自行修改的结构体中
type ReqShipperDelete struct {
	Id int `json:"id"`
}
