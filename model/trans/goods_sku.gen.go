package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsSkuList 你可以把ReqGoodsSkuList嵌套到需要自行修改的结构体中
type ReqGoodsSkuList struct {
	ReqPage
	mysql.GoodsSku
}

// ReqGoodsSkuCreate 你可以把ReqGoodsSkuCreate或mysql.GoodsSku嵌套到需要自行修改的结构体中
type ReqGoodsSkuCreate = mysql.GoodsSku

// ReqGoodsSkuUpdate 你可以把ReqGoodsSkuUpdate或mysql.GoodsSku嵌套到需要自行修改的结构体中
type ReqGoodsSkuUpdate = mysql.GoodsSku

// ReqGoodsSkuDelete 你可以把ReqGoodsSkuDelete嵌套到需要自行修改的结构体中
type ReqGoodsSkuDelete struct {
	Id int `json:"id"`
}
