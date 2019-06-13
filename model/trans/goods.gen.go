package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsList 你可以把ReqGoodsList嵌套到需要自行修改的结构体中
type ReqGoodsList struct {
	ReqPage
	mysql.Goods
}

// ReqGoodsCreate 你可以把ReqGoodsCreate或mysql.Goods嵌套到需要自行修改的结构体中
type ReqGoodsCreate = mysql.Goods

// ReqGoodsUpdate 你可以把ReqGoodsUpdate或mysql.Goods嵌套到需要自行修改的结构体中
type ReqGoodsUpdate = mysql.Goods

// ReqGoodsDelete 你可以把ReqGoodsDelete嵌套到需要自行修改的结构体中
type ReqGoodsDelete struct {
	Id int `json:"id"`
}
