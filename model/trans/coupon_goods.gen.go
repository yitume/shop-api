package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqCouponGoodsList 你可以把ReqCouponGoodsList嵌套到需要自行修改的结构体中
type ReqCouponGoodsList struct {
	ReqPage
	mysql.CouponGoods
}

// ReqCouponGoodsCreate 你可以把ReqCouponGoodsCreate或mysql.CouponGoods嵌套到需要自行修改的结构体中
type ReqCouponGoodsCreate = mysql.CouponGoods

// ReqCouponGoodsUpdate 你可以把ReqCouponGoodsUpdate或mysql.CouponGoods嵌套到需要自行修改的结构体中
type ReqCouponGoodsUpdate = mysql.CouponGoods

// ReqCouponGoodsDelete 你可以把ReqCouponGoodsDelete嵌套到需要自行修改的结构体中
type ReqCouponGoodsDelete struct {
	Id int `json:"id"`
}
