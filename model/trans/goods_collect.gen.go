package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsCollectList 你可以把ReqGoodsCollectList嵌套到需要自行修改的结构体中
type ReqGoodsCollectList struct {
	ReqPage
	mysql.GoodsCollect
}

// ReqGoodsCollectCreate 你可以把ReqGoodsCollectCreate或mysql.GoodsCollect嵌套到需要自行修改的结构体中
type ReqGoodsCollectCreate = mysql.GoodsCollect

// ReqGoodsCollectUpdate 你可以把ReqGoodsCollectUpdate或mysql.GoodsCollect嵌套到需要自行修改的结构体中
type ReqGoodsCollectUpdate = mysql.GoodsCollect

// ReqGoodsCollectDelete 你可以把ReqGoodsCollectDelete嵌套到需要自行修改的结构体中
type ReqGoodsCollectDelete struct {
	Id int `json:"id"`
}
