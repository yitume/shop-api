package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqVersionList 你可以把ReqVersionList嵌套到需要自行修改的结构体中
type ReqVersionList struct {
	ReqPage
	mysql.Version
}

// ReqVersionCreate 你可以把ReqVersionCreate或mysql.Version嵌套到需要自行修改的结构体中
type ReqVersionCreate = mysql.Version

// ReqVersionUpdate 你可以把ReqVersionUpdate或mysql.Version嵌套到需要自行修改的结构体中
type ReqVersionUpdate = mysql.Version

// ReqVersionDelete 你可以把ReqVersionDelete嵌套到需要自行修改的结构体中
type ReqVersionDelete struct {
	Id int `json:"id"`
}
