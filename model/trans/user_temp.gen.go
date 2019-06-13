package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserTempList 你可以把ReqUserTempList嵌套到需要自行修改的结构体中
type ReqUserTempList struct {
	ReqPage
	mysql.UserTemp
}

// ReqUserTempCreate 你可以把ReqUserTempCreate或mysql.UserTemp嵌套到需要自行修改的结构体中
type ReqUserTempCreate = mysql.UserTemp

// ReqUserTempUpdate 你可以把ReqUserTempUpdate或mysql.UserTemp嵌套到需要自行修改的结构体中
type ReqUserTempUpdate = mysql.UserTemp

// ReqUserTempDelete 你可以把ReqUserTempDelete嵌套到需要自行修改的结构体中
type ReqUserTempDelete struct {
	UserId int `json:"user_id"`
}
