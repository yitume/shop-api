package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqAuthRuleList 你可以把ReqAuthRuleList嵌套到需要自行修改的结构体中
type ReqAuthRuleList struct {
	ReqPage
	mysql.AuthRule
}

// ReqAuthRuleCreate 你可以把ReqAuthRuleCreate或mysql.AuthRule嵌套到需要自行修改的结构体中
type ReqAuthRuleCreate = mysql.AuthRule

// ReqAuthRuleUpdate 你可以把ReqAuthRuleUpdate或mysql.AuthRule嵌套到需要自行修改的结构体中
type ReqAuthRuleUpdate = mysql.AuthRule

// ReqAuthRuleDelete 你可以把ReqAuthRuleDelete嵌套到需要自行修改的结构体中
type ReqAuthRuleDelete struct {
	Id int `json:"id"`
}
