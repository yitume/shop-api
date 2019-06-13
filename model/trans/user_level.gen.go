package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserLevelList 你可以把ReqUserLevelList嵌套到需要自行修改的结构体中
type ReqUserLevelList struct {
	ReqPage
	mysql.UserLevel
}

// ReqUserLevelCreate 你可以把ReqUserLevelCreate或mysql.UserLevel嵌套到需要自行修改的结构体中
type ReqUserLevelCreate = mysql.UserLevel

// ReqUserLevelUpdate 你可以把ReqUserLevelUpdate或mysql.UserLevel嵌套到需要自行修改的结构体中
type ReqUserLevelUpdate = mysql.UserLevel

// ReqUserLevelDelete 你可以把ReqUserLevelDelete嵌套到需要自行修改的结构体中
type ReqUserLevelDelete struct {
	Id int `json:"id"`
}
