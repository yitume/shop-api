package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqAuthGroupAccessList 你可以把ReqAuthGroupAccessList嵌套到需要自行修改的结构体中
type ReqAuthGroupAccessList struct {
	ReqPage
	mysql.AuthGroupAccess
}

// ReqAuthGroupAccessCreate 你可以把ReqAuthGroupAccessCreate或mysql.AuthGroupAccess嵌套到需要自行修改的结构体中
type ReqAuthGroupAccessCreate = mysql.AuthGroupAccess

// ReqAuthGroupAccessUpdate 你可以把ReqAuthGroupAccessUpdate或mysql.AuthGroupAccess嵌套到需要自行修改的结构体中
type ReqAuthGroupAccessUpdate = mysql.AuthGroupAccess

// ReqAuthGroupAccessDelete 你可以把ReqAuthGroupAccessDelete嵌套到需要自行修改的结构体中
type ReqAuthGroupAccessDelete struct {
	Uid     int `json:"uid"`
	GroupId int `json:"group_id"`
}
