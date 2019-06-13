package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqDiscountGoodsList 你可以把ReqDiscountGoodsList嵌套到需要自行修改的结构体中
type ReqDiscountGoodsList struct {
	ReqPage
	mysql.DiscountGoods
}

// ReqDiscountGoodsCreate 你可以把ReqDiscountGoodsCreate或mysql.DiscountGoods嵌套到需要自行修改的结构体中
type ReqDiscountGoodsCreate = mysql.DiscountGoods

// ReqDiscountGoodsUpdate 你可以把ReqDiscountGoodsUpdate或mysql.DiscountGoods嵌套到需要自行修改的结构体中
type ReqDiscountGoodsUpdate = mysql.DiscountGoods

// ReqDiscountGoodsDelete 你可以把ReqDiscountGoodsDelete嵌套到需要自行修改的结构体中
type ReqDiscountGoodsDelete struct {
	Id int `json:"id"`
}
