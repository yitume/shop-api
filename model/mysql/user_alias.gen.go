package mysql

type UserAlias struct {
	Id          int   `json:"id" form:"id" gorm:"primary_key"`     // ID
	UserId      int   `json:"user_id" form:"user_id" `             // 主账号ID[user_id]
	AliasUserId int   `json:"alias_user_id" form:"alias_user_id" ` // 第三方帐号ID[user_id 占位行的id]
	DeleteTime  int64 `json:"delete_time" form:"delete_time" `     // 软删除时间
	OpenId      int   `json:"open_id" form:"open_id" `             // 商家ID

}

func (*UserAlias) TableName() string {
	return "user_alias"
}
