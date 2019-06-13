package mysql

type UserLevel struct {
	Id          int    `json:"id" form:"id" gorm:"primary_key"`   // id
	Title       string `json:"title" form:"title" `               // 级别名称 称号
	GrowthValue int    `json:"growth_value" form:"growth_value" ` // 成长值 经验
	Desc        string `json:"desc" form:"desc" `                 // 描述   福利
	DeleteTime  int64  `json:"delete_time" form:"delete_time" `   // 软删除时间
	OpenId      int    `json:"open_id" form:"open_id" `           // 商家ID

}

func (*UserLevel) TableName() string {
	return "user_level"
}
