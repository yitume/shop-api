package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserAssetsList 你可以把ReqUserAssetsList嵌套到需要自行修改的结构体中
type ReqUserAssetsList struct {
	ReqPage
	mysql.UserAssets
}

// ReqUserAssetsCreate 你可以把ReqUserAssetsCreate或mysql.UserAssets嵌套到需要自行修改的结构体中
type ReqUserAssetsCreate = mysql.UserAssets

// ReqUserAssetsUpdate 你可以把ReqUserAssetsUpdate或mysql.UserAssets嵌套到需要自行修改的结构体中
type ReqUserAssetsUpdate = mysql.UserAssets

// ReqUserAssetsDelete 你可以把ReqUserAssetsDelete嵌套到需要自行修改的结构体中
type ReqUserAssetsDelete struct {
	Id int `json:"id"`
}
