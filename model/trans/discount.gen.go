package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqDiscountList 你可以把ReqDiscountList嵌套到需要自行修改的结构体中
type ReqDiscountList struct {
	ReqPage
	mysql.Discount
}

// ReqDiscountCreate 你可以把ReqDiscountCreate或mysql.Discount嵌套到需要自行修改的结构体中
type ReqDiscountCreate = mysql.Discount

// ReqDiscountUpdate 你可以把ReqDiscountUpdate或mysql.Discount嵌套到需要自行修改的结构体中
type ReqDiscountUpdate = mysql.Discount

// ReqDiscountDelete 你可以把ReqDiscountDelete嵌套到需要自行修改的结构体中
type ReqDiscountDelete struct {
	Id int `json:"id"`
}
