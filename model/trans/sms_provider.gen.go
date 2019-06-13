package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqSmsProviderList 你可以把ReqSmsProviderList嵌套到需要自行修改的结构体中
type ReqSmsProviderList struct {
	ReqPage
	mysql.SmsProvider
}

// ReqSmsProviderCreate 你可以把ReqSmsProviderCreate或mysql.SmsProvider嵌套到需要自行修改的结构体中
type ReqSmsProviderCreate = mysql.SmsProvider

// ReqSmsProviderUpdate 你可以把ReqSmsProviderUpdate或mysql.SmsProvider嵌套到需要自行修改的结构体中
type ReqSmsProviderUpdate = mysql.SmsProvider

// ReqSmsProviderDelete 你可以把ReqSmsProviderDelete嵌套到需要自行修改的结构体中
type ReqSmsProviderDelete struct {
	Type string `json:"type"`
}
