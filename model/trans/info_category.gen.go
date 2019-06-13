package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqInfoCategoryList 你可以把ReqInfoCategoryList嵌套到需要自行修改的结构体中
type ReqInfoCategoryList struct {
	ReqPage
	mysql.InfoCategory
}

// ReqInfoCategoryCreate 你可以把ReqInfoCategoryCreate或mysql.InfoCategory嵌套到需要自行修改的结构体中
type ReqInfoCategoryCreate = mysql.InfoCategory

// ReqInfoCategoryUpdate 你可以把ReqInfoCategoryUpdate或mysql.InfoCategory嵌套到需要自行修改的结构体中
type ReqInfoCategoryUpdate = mysql.InfoCategory

// ReqInfoCategoryDelete 你可以把ReqInfoCategoryDelete嵌套到需要自行修改的结构体中
type ReqInfoCategoryDelete struct {
	Id int `json:"id"`
}
