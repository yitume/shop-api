package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqFullcutList 你可以把ReqFullcutList嵌套到需要自行修改的结构体中
type ReqFullcutList struct {
	ReqPage
	mysql.Fullcut
}

// ReqFullcutCreate 你可以把ReqFullcutCreate或mysql.Fullcut嵌套到需要自行修改的结构体中
type ReqFullcutCreate = mysql.Fullcut

// ReqFullcutUpdate 你可以把ReqFullcutUpdate或mysql.Fullcut嵌套到需要自行修改的结构体中
type ReqFullcutUpdate = mysql.Fullcut

// ReqFullcutDelete 你可以把ReqFullcutDelete嵌套到需要自行修改的结构体中
type ReqFullcutDelete struct {
	Id int `json:"id"`
}
