package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsEvaluateList 你可以把ReqGoodsEvaluateList嵌套到需要自行修改的结构体中
type ReqGoodsEvaluateList struct {
	ReqPage
	mysql.GoodsEvaluate
}

// ReqGoodsEvaluateCreate 你可以把ReqGoodsEvaluateCreate或mysql.GoodsEvaluate嵌套到需要自行修改的结构体中
type ReqGoodsEvaluateCreate = mysql.GoodsEvaluate

// ReqGoodsEvaluateUpdate 你可以把ReqGoodsEvaluateUpdate或mysql.GoodsEvaluate嵌套到需要自行修改的结构体中
type ReqGoodsEvaluateUpdate = mysql.GoodsEvaluate

// ReqGoodsEvaluateDelete 你可以把ReqGoodsEvaluateDelete嵌套到需要自行修改的结构体中
type ReqGoodsEvaluateDelete struct {
	Id int `json:"id"`
}
