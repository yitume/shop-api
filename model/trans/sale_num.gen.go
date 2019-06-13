package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqSaleNumList 你可以把ReqSaleNumList嵌套到需要自行修改的结构体中
type ReqSaleNumList struct {
	ReqPage
	mysql.SaleNum
}

// ReqSaleNumCreate 你可以把ReqSaleNumCreate或mysql.SaleNum嵌套到需要自行修改的结构体中
type ReqSaleNumCreate = mysql.SaleNum

// ReqSaleNumUpdate 你可以把ReqSaleNumUpdate或mysql.SaleNum嵌套到需要自行修改的结构体中
type ReqSaleNumUpdate = mysql.SaleNum

// ReqSaleNumDelete 你可以把ReqSaleNumDelete嵌套到需要自行修改的结构体中
type ReqSaleNumDelete struct {
}
