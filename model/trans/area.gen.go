package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqAreaList 你可以把ReqAreaList嵌套到需要自行修改的结构体中
type ReqAreaList struct {
	ReqPage
	mysql.Area
}

// ReqAreaCreate 你可以把ReqAreaCreate或mysql.Area嵌套到需要自行修改的结构体中
type ReqAreaCreate = mysql.Area

// ReqAreaUpdate 你可以把ReqAreaUpdate或mysql.Area嵌套到需要自行修改的结构体中
type ReqAreaUpdate = mysql.Area

// ReqAreaDelete 你可以把ReqAreaDelete嵌套到需要自行修改的结构体中
type ReqAreaDelete struct {
	Id int `json:"id"`
}
