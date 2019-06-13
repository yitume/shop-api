package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserOpenList 你可以把ReqUserOpenList嵌套到需要自行修改的结构体中
type ReqUserOpenList struct {
	ReqPage
	mysql.UserOpen
}

// ReqUserOpenCreate 你可以把ReqUserOpenCreate或mysql.UserOpen嵌套到需要自行修改的结构体中
type ReqUserOpenCreate = mysql.UserOpen

// ReqUserOpenUpdate 你可以把ReqUserOpenUpdate或mysql.UserOpen嵌套到需要自行修改的结构体中
type ReqUserOpenUpdate = mysql.UserOpen

// ReqUserOpenDelete 你可以把ReqUserOpenDelete嵌套到需要自行修改的结构体中
type ReqUserOpenDelete struct {
	Uid int `json:"uid"`
}
