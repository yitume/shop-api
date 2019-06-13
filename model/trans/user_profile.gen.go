package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqUserProfileList 你可以把ReqUserProfileList嵌套到需要自行修改的结构体中
type ReqUserProfileList struct {
	ReqPage
	mysql.UserProfile
}

// ReqUserProfileCreate 你可以把ReqUserProfileCreate或mysql.UserProfile嵌套到需要自行修改的结构体中
type ReqUserProfileCreate = mysql.UserProfile

// ReqUserProfileUpdate 你可以把ReqUserProfileUpdate或mysql.UserProfile嵌套到需要自行修改的结构体中
type ReqUserProfileUpdate = mysql.UserProfile

// ReqUserProfileDelete 你可以把ReqUserProfileDelete嵌套到需要自行修改的结构体中
type ReqUserProfileDelete struct {
	Id int `json:"id"`
}
