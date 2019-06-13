package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqSmsSceneList 你可以把ReqSmsSceneList嵌套到需要自行修改的结构体中
type ReqSmsSceneList struct {
	ReqPage
	mysql.SmsScene
}

// ReqSmsSceneCreate 你可以把ReqSmsSceneCreate或mysql.SmsScene嵌套到需要自行修改的结构体中
type ReqSmsSceneCreate = mysql.SmsScene

// ReqSmsSceneUpdate 你可以把ReqSmsSceneUpdate或mysql.SmsScene嵌套到需要自行修改的结构体中
type ReqSmsSceneUpdate = mysql.SmsScene

// ReqSmsSceneDelete 你可以把ReqSmsSceneDelete嵌套到需要自行修改的结构体中
type ReqSmsSceneDelete struct {
	Id int `json:"id"`
}
