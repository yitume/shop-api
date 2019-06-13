package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqSmsList 你可以把ReqSmsList嵌套到需要自行修改的结构体中
type ReqSmsList struct {
	ReqPage
	mysql.Sms
}

// ReqSmsCreate 你可以把ReqSmsCreate或mysql.Sms嵌套到需要自行修改的结构体中
type ReqSmsCreate = mysql.Sms

// ReqSmsUpdate 你可以把ReqSmsUpdate或mysql.Sms嵌套到需要自行修改的结构体中
type ReqSmsUpdate = mysql.Sms

// ReqSmsDelete 你可以把ReqSmsDelete嵌套到需要自行修改的结构体中
type ReqSmsDelete struct {
	Id int `json:"id"`
}
