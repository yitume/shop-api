package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGroupGoodsList 你可以把ReqGroupGoodsList嵌套到需要自行修改的结构体中
type ReqGroupGoodsList struct {
	ReqPage
	mysql.GroupGoods
}

// ReqGroupGoodsCreate 你可以把ReqGroupGoodsCreate或mysql.GroupGoods嵌套到需要自行修改的结构体中
type ReqGroupGoodsCreate = mysql.GroupGoods

// ReqGroupGoodsUpdate 你可以把ReqGroupGoodsUpdate或mysql.GroupGoods嵌套到需要自行修改的结构体中
type ReqGroupGoodsUpdate = mysql.GroupGoods

// ReqGroupGoodsDelete 你可以把ReqGroupGoodsDelete嵌套到需要自行修改的结构体中
type ReqGroupGoodsDelete struct {
	Id int `json:"id"`
}
