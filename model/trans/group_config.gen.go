package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGroupConfigList 你可以把ReqGroupConfigList嵌套到需要自行修改的结构体中
type ReqGroupConfigList struct {
	ReqPage
	mysql.GroupConfig
}

// ReqGroupConfigCreate 你可以把ReqGroupConfigCreate或mysql.GroupConfig嵌套到需要自行修改的结构体中
type ReqGroupConfigCreate = mysql.GroupConfig

// ReqGroupConfigUpdate 你可以把ReqGroupConfigUpdate或mysql.GroupConfig嵌套到需要自行修改的结构体中
type ReqGroupConfigUpdate = mysql.GroupConfig

// ReqGroupConfigDelete 你可以把ReqGroupConfigDelete嵌套到需要自行修改的结构体中
type ReqGroupConfigDelete struct {
}
