package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqMaterialList 你可以把ReqMaterialList嵌套到需要自行修改的结构体中
type ReqMaterialList struct {
	ReqPage
	mysql.Material
}

// ReqMaterialCreate 你可以把ReqMaterialCreate或mysql.Material嵌套到需要自行修改的结构体中
type ReqMaterialCreate = mysql.Material

// ReqMaterialUpdate 你可以把ReqMaterialUpdate或mysql.Material嵌套到需要自行修改的结构体中
type ReqMaterialUpdate = mysql.Material

// ReqMaterialDelete 你可以把ReqMaterialDelete嵌套到需要自行修改的结构体中
type ReqMaterialDelete struct {
	Id int `json:"id"`
}
