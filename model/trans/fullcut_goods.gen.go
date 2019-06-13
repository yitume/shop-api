package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqFullcutGoodsList 你可以把ReqFullcutGoodsList嵌套到需要自行修改的结构体中
type ReqFullcutGoodsList struct {
	ReqPage
	mysql.FullcutGoods
}

// ReqFullcutGoodsCreate 你可以把ReqFullcutGoodsCreate或mysql.FullcutGoods嵌套到需要自行修改的结构体中
type ReqFullcutGoodsCreate = mysql.FullcutGoods

// ReqFullcutGoodsUpdate 你可以把ReqFullcutGoodsUpdate或mysql.FullcutGoods嵌套到需要自行修改的结构体中
type ReqFullcutGoodsUpdate = mysql.FullcutGoods

// ReqFullcutGoodsDelete 你可以把ReqFullcutGoodsDelete嵌套到需要自行修改的结构体中
type ReqFullcutGoodsDelete struct {
	Id int `json:"id"`
}
