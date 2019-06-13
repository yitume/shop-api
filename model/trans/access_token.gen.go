package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqAccessTokenList 你可以把ReqAccessTokenList嵌套到需要自行修改的结构体中
type ReqAccessTokenList struct {
	ReqPage
	mysql.AccessToken
}

// ReqAccessTokenCreate 你可以把ReqAccessTokenCreate或mysql.AccessToken嵌套到需要自行修改的结构体中
type ReqAccessTokenCreate = mysql.AccessToken

// ReqAccessTokenUpdate 你可以把ReqAccessTokenUpdate或mysql.AccessToken嵌套到需要自行修改的结构体中
type ReqAccessTokenUpdate = mysql.AccessToken

// ReqAccessTokenDelete 你可以把ReqAccessTokenDelete嵌套到需要自行修改的结构体中
type ReqAccessTokenDelete struct {
	Jti int `json:"jti"`
}
