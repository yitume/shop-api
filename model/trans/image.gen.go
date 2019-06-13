package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqImageList 你可以把ReqImageList嵌套到需要自行修改的结构体中
type ReqImageList struct {
	ReqPage
	mysql.Image
}

// ReqImageCreate 你可以把ReqImageCreate或mysql.Image嵌套到需要自行修改的结构体中
type ReqImageCreate = mysql.Image

// ReqImageUpdate 你可以把ReqImageUpdate或mysql.Image嵌套到需要自行修改的结构体中
type ReqImageUpdate = mysql.Image

// ReqImageDelete 你可以把ReqImageDelete嵌套到需要自行修改的结构体中
type ReqImageDelete struct {
	Id int `json:"id"`
}
