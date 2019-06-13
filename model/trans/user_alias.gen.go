package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserAliasList 你可以把ReqUserAliasList嵌套到需要自行修改的结构体中
type ReqUserAliasList struct {
	ReqPage
	mysql.UserAlias
}

// ReqUserAliasCreate 你可以把ReqUserAliasCreate或mysql.UserAlias嵌套到需要自行修改的结构体中
type ReqUserAliasCreate = mysql.UserAlias

// ReqUserAliasUpdate 你可以把ReqUserAliasUpdate或mysql.UserAlias嵌套到需要自行修改的结构体中
type ReqUserAliasUpdate = mysql.UserAlias

// ReqUserAliasDelete 你可以把ReqUserAliasDelete嵌套到需要自行修改的结构体中
type ReqUserAliasDelete struct {
	Id int `json:"id"`
}
