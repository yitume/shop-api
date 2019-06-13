package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUploadList 你可以把ReqUploadList嵌套到需要自行修改的结构体中
type ReqUploadList struct {
	ReqPage
	mysql.Upload
}

// ReqUploadCreate 你可以把ReqUploadCreate或mysql.Upload嵌套到需要自行修改的结构体中
type ReqUploadCreate = mysql.Upload

// ReqUploadUpdate 你可以把ReqUploadUpdate或mysql.Upload嵌套到需要自行修改的结构体中
type ReqUploadUpdate = mysql.Upload

// ReqUploadDelete 你可以把ReqUploadDelete嵌套到需要自行修改的结构体中
type ReqUploadDelete struct {
	Id int `json:"id"`
}
