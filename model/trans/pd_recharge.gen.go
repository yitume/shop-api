package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqPdRechargeList 你可以把ReqPdRechargeList嵌套到需要自行修改的结构体中
type ReqPdRechargeList struct {
	ReqPage
	mysql.PdRecharge
}

// ReqPdRechargeCreate 你可以把ReqPdRechargeCreate或mysql.PdRecharge嵌套到需要自行修改的结构体中
type ReqPdRechargeCreate = mysql.PdRecharge

// ReqPdRechargeUpdate 你可以把ReqPdRechargeUpdate或mysql.PdRecharge嵌套到需要自行修改的结构体中
type ReqPdRechargeUpdate = mysql.PdRecharge

// ReqPdRechargeDelete 你可以把ReqPdRechargeDelete嵌套到需要自行修改的结构体中
type ReqPdRechargeDelete struct {
	Id int `json:"id"`
}
