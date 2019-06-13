package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqOrderGoodsList 你可以把ReqOrderGoodsList嵌套到需要自行修改的结构体中
type ReqOrderGoodsList struct {
	ReqPage
	mysql.OrderGoods
}

// ReqOrderGoodsCreate 你可以把ReqOrderGoodsCreate或mysql.OrderGoods嵌套到需要自行修改的结构体中
type ReqOrderGoodsCreate = mysql.OrderGoods

// ReqOrderGoodsUpdate 你可以把ReqOrderGoodsUpdate或mysql.OrderGoods嵌套到需要自行修改的结构体中
type ReqOrderGoodsUpdate = mysql.OrderGoods

// ReqOrderGoodsDelete 你可以把ReqOrderGoodsDelete嵌套到需要自行修改的结构体中
type ReqOrderGoodsDelete struct {
	Id int `json:"id"`
}
