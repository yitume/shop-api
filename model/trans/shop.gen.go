package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqShopList 你可以把ReqShopList嵌套到需要自行修改的结构体中
type ReqShopList struct {
	ReqPage
	mysql.Shop
}

// ReqShopCreate 你可以把ReqShopCreate或mysql.Shop嵌套到需要自行修改的结构体中
type ReqShopCreate = mysql.Shop

// ReqShopUpdate 你可以把ReqShopUpdate或mysql.Shop嵌套到需要自行修改的结构体中
type ReqShopUpdate = mysql.Shop

// ReqShopDelete 你可以把ReqShopDelete嵌套到需要自行修改的结构体中
type ReqShopDelete struct {
	OpenId int `json:"open_id"`
}
