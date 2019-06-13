package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqCouponList 你可以把ReqCouponList嵌套到需要自行修改的结构体中
type ReqCouponList struct {
	ReqPage
	mysql.Coupon
}

// ReqCouponCreate 你可以把ReqCouponCreate或mysql.Coupon嵌套到需要自行修改的结构体中
type ReqCouponCreate = mysql.Coupon

// ReqCouponUpdate 你可以把ReqCouponUpdate或mysql.Coupon嵌套到需要自行修改的结构体中
type ReqCouponUpdate = mysql.Coupon

// ReqCouponDelete 你可以把ReqCouponDelete嵌套到需要自行修改的结构体中
type ReqCouponDelete struct {
	Id int `json:"id"`
}
