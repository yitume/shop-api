package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGroupList 你可以把ReqGroupList嵌套到需要自行修改的结构体中
type ReqGroupList struct {
	ReqPage
	mysql.Group
}

// ReqGroupCreate 你可以把ReqGroupCreate或mysql.Group嵌套到需要自行修改的结构体中
type ReqGroupCreate = mysql.Group

// ReqGroupUpdate 你可以把ReqGroupUpdate或mysql.Group嵌套到需要自行修改的结构体中
type ReqGroupUpdate = mysql.Group

// ReqGroupDelete 你可以把ReqGroupDelete嵌套到需要自行修改的结构体中
type ReqGroupDelete struct {
	Id int `json:"id"`
}
