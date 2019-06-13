package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqPdCashList 你可以把ReqPdCashList嵌套到需要自行修改的结构体中
type ReqPdCashList struct {
	ReqPage
	mysql.PdCash
}

// ReqPdCashCreate 你可以把ReqPdCashCreate或mysql.PdCash嵌套到需要自行修改的结构体中
type ReqPdCashCreate = mysql.PdCash

// ReqPdCashUpdate 你可以把ReqPdCashUpdate或mysql.PdCash嵌套到需要自行修改的结构体中
type ReqPdCashUpdate = mysql.PdCash

// ReqPdCashDelete 你可以把ReqPdCashDelete嵌套到需要自行修改的结构体中
type ReqPdCashDelete struct {
	Id int `json:"id"`
}
