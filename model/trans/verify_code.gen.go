package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqVerifyCodeList 你可以把ReqVerifyCodeList嵌套到需要自行修改的结构体中
type ReqVerifyCodeList struct {
	ReqPage
	mysql.VerifyCode
}

// ReqVerifyCodeCreate 你可以把ReqVerifyCodeCreate或mysql.VerifyCode嵌套到需要自行修改的结构体中
type ReqVerifyCodeCreate = mysql.VerifyCode

// ReqVerifyCodeUpdate 你可以把ReqVerifyCodeUpdate或mysql.VerifyCode嵌套到需要自行修改的结构体中
type ReqVerifyCodeUpdate = mysql.VerifyCode

// ReqVerifyCodeDelete 你可以把ReqVerifyCodeDelete嵌套到需要自行修改的结构体中
type ReqVerifyCodeDelete struct {
	Id int `json:"id"`
}
