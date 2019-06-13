package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsSpecList 你可以把ReqGoodsSpecList嵌套到需要自行修改的结构体中
type ReqGoodsSpecList struct {
	ReqPage
	mysql.GoodsSpec
}

// ReqGoodsSpecCreate 你可以把ReqGoodsSpecCreate或mysql.GoodsSpec嵌套到需要自行修改的结构体中
type ReqGoodsSpecCreate = mysql.GoodsSpec

// ReqGoodsSpecUpdate 你可以把ReqGoodsSpecUpdate或mysql.GoodsSpec嵌套到需要自行修改的结构体中
type ReqGoodsSpecUpdate = mysql.GoodsSpec

// ReqGoodsSpecDelete 你可以把ReqGoodsSpecDelete嵌套到需要自行修改的结构体中
type ReqGoodsSpecDelete struct {
	Id int `json:"id"`
}
