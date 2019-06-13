package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqGoodsImageList 你可以把ReqGoodsImageList嵌套到需要自行修改的结构体中
type ReqGoodsImageList struct {
	ReqPage
	mysql.GoodsImage
}

// ReqGoodsImageCreate 你可以把ReqGoodsImageCreate或mysql.GoodsImage嵌套到需要自行修改的结构体中
type ReqGoodsImageCreate = mysql.GoodsImage

// ReqGoodsImageUpdate 你可以把ReqGoodsImageUpdate或mysql.GoodsImage嵌套到需要自行修改的结构体中
type ReqGoodsImageUpdate = mysql.GoodsImage

// ReqGoodsImageDelete 你可以把ReqGoodsImageDelete嵌套到需要自行修改的结构体中
type ReqGoodsImageDelete struct {
	Id int `json:"id"`
}
