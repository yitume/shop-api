package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqAuthGroupList 你可以把ReqAuthGroupList嵌套到需要自行修改的结构体中
type ReqAuthGroupList struct {
	ReqPage
	mysql.AuthGroup
}

// ReqAuthGroupCreate 你可以把ReqAuthGroupCreate或mysql.AuthGroup嵌套到需要自行修改的结构体中
type ReqAuthGroupCreate = mysql.AuthGroup

// ReqAuthGroupUpdate 你可以把ReqAuthGroupUpdate或mysql.AuthGroup嵌套到需要自行修改的结构体中
type ReqAuthGroupUpdate = mysql.AuthGroup

// ReqAuthGroupDelete 你可以把ReqAuthGroupDelete嵌套到需要自行修改的结构体中
type ReqAuthGroupDelete struct {
	Id int `json:"id"`
}
