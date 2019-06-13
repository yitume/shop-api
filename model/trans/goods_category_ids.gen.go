package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsCategoryIdsList 你可以把ReqGoodsCategoryIdsList嵌套到需要自行修改的结构体中
type ReqGoodsCategoryIdsList struct {
	ReqPage
	mysql.GoodsCategoryIds
}

// ReqGoodsCategoryIdsCreate 你可以把ReqGoodsCategoryIdsCreate或mysql.GoodsCategoryIds嵌套到需要自行修改的结构体中
type ReqGoodsCategoryIdsCreate = mysql.GoodsCategoryIds

// ReqGoodsCategoryIdsUpdate 你可以把ReqGoodsCategoryIdsUpdate或mysql.GoodsCategoryIds嵌套到需要自行修改的结构体中
type ReqGoodsCategoryIdsUpdate = mysql.GoodsCategoryIds

// ReqGoodsCategoryIdsDelete 你可以把ReqGoodsCategoryIdsDelete嵌套到需要自行修改的结构体中
type ReqGoodsCategoryIdsDelete struct {
	Id int `json:"id"`
}
