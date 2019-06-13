package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserAccountList 你可以把ReqUserAccountList嵌套到需要自行修改的结构体中
type ReqUserAccountList struct {
	ReqPage
	mysql.UserAccount
}

// ReqUserAccountCreate 你可以把ReqUserAccountCreate或mysql.UserAccount嵌套到需要自行修改的结构体中
type ReqUserAccountCreate = mysql.UserAccount

// ReqUserAccountUpdate 你可以把ReqUserAccountUpdate或mysql.UserAccount嵌套到需要自行修改的结构体中
type ReqUserAccountUpdate = mysql.UserAccount

// ReqUserAccountDelete 你可以把ReqUserAccountDelete嵌套到需要自行修改的结构体中
type ReqUserAccountDelete struct {
	Uid int `json:"uid"`
}
