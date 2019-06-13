package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsCategoryList 你可以把ReqGoodsCategoryList嵌套到需要自行修改的结构体中
type ReqGoodsCategoryList struct {
	ReqPage
	mysql.GoodsCategory
}

// ReqGoodsCategoryCreate 你可以把ReqGoodsCategoryCreate或mysql.GoodsCategory嵌套到需要自行修改的结构体中
type ReqGoodsCategoryCreate = mysql.GoodsCategory

// ReqGoodsCategoryUpdate 你可以把ReqGoodsCategoryUpdate或mysql.GoodsCategory嵌套到需要自行修改的结构体中
type ReqGoodsCategoryUpdate = mysql.GoodsCategory

// ReqGoodsCategoryDelete 你可以把ReqGoodsCategoryDelete嵌套到需要自行修改的结构体中
type ReqGoodsCategoryDelete struct {
	Id int `json:"id"`
}
