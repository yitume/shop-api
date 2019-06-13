package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqBizList 你可以把ReqBizList嵌套到需要自行修改的结构体中
type ReqBizList struct {
	ReqPage
	mysql.Biz
}

// ReqBizCreate 你可以把ReqBizCreate或mysql.Biz嵌套到需要自行修改的结构体中
type ReqBizCreate = mysql.Biz

// ReqBizUpdate 你可以把ReqBizUpdate或mysql.Biz嵌套到需要自行修改的结构体中
type ReqBizUpdate = mysql.Biz

// ReqBizDelete 你可以把ReqBizDelete嵌套到需要自行修改的结构体中
type ReqBizDelete struct {
	OpenId int `json:"open_id"`
}
