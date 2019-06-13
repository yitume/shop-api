package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqAddressList 你可以把ReqAddressList嵌套到需要自行修改的结构体中
type ReqAddressList struct {
	ReqPage
	mysql.Address
}

// ReqAddressCreate 你可以把ReqAddressCreate或mysql.Address嵌套到需要自行修改的结构体中
type ReqAddressCreate = mysql.Address

// ReqAddressUpdate 你可以把ReqAddressUpdate或mysql.Address嵌套到需要自行修改的结构体中
type ReqAddressUpdate = mysql.Address

// ReqAddressDelete 你可以把ReqAddressDelete嵌套到需要自行修改的结构体中
type ReqAddressDelete struct {
	Id int `json:"id"`
}
