package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsSpecValueList 你可以把ReqGoodsSpecValueList嵌套到需要自行修改的结构体中
type ReqGoodsSpecValueList struct {
	ReqPage
	mysql.GoodsSpecValue
}

// ReqGoodsSpecValueCreate 你可以把ReqGoodsSpecValueCreate或mysql.GoodsSpecValue嵌套到需要自行修改的结构体中
type ReqGoodsSpecValueCreate = mysql.GoodsSpecValue

// ReqGoodsSpecValueUpdate 你可以把ReqGoodsSpecValueUpdate或mysql.GoodsSpecValue嵌套到需要自行修改的结构体中
type ReqGoodsSpecValueUpdate = mysql.GoodsSpecValue

// ReqGoodsSpecValueDelete 你可以把ReqGoodsSpecValueDelete嵌套到需要自行修改的结构体中
type ReqGoodsSpecValueDelete struct {
	Id int `json:"id"`
}
