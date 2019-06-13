package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqPluginList 你可以把ReqPluginList嵌套到需要自行修改的结构体中
type ReqPluginList struct {
	ReqPage
	mysql.Plugin
}

// ReqPluginCreate 你可以把ReqPluginCreate或mysql.Plugin嵌套到需要自行修改的结构体中
type ReqPluginCreate = mysql.Plugin

// ReqPluginUpdate 你可以把ReqPluginUpdate或mysql.Plugin嵌套到需要自行修改的结构体中
type ReqPluginUpdate = mysql.Plugin

// ReqPluginDelete 你可以把ReqPluginDelete嵌套到需要自行修改的结构体中
type ReqPluginDelete struct {
}
