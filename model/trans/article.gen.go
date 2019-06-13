package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqArticleList 你可以把ReqArticleList嵌套到需要自行修改的结构体中
type ReqArticleList struct {
	ReqPage
	mysql.Article
}

// ReqArticleCreate 你可以把ReqArticleCreate或mysql.Article嵌套到需要自行修改的结构体中
type ReqArticleCreate = mysql.Article

// ReqArticleUpdate 你可以把ReqArticleUpdate或mysql.Article嵌套到需要自行修改的结构体中
type ReqArticleUpdate = mysql.Article

// ReqArticleDelete 你可以把ReqArticleDelete嵌套到需要自行修改的结构体中
type ReqArticleDelete struct {
	Id int `json:"id"`
}
