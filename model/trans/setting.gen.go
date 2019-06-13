package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqSettingList 你可以把ReqSettingList嵌套到需要自行修改的结构体中
type ReqSettingList struct {
	ReqPage
	mysql.Setting
}

// ReqSettingCreate 你可以把ReqSettingCreate或mysql.Setting嵌套到需要自行修改的结构体中
type ReqSettingCreate = mysql.Setting

// ReqSettingUpdate 你可以把ReqSettingUpdate或mysql.Setting嵌套到需要自行修改的结构体中
type ReqSettingUpdate = mysql.Setting

// ReqSettingDelete 你可以把ReqSettingDelete嵌套到需要自行修改的结构体中
type ReqSettingDelete struct {
	Key string `json:"key"`
}
