package mysql

type Visit struct {
	Id              int    `json:"id" form:"id" gorm:"primary_key"`             // 主键
	UserId          int    `json:"user_id" form:"user_id" `                     // 使用UID
	ModelRelationId int    `json:"model_relation_id" form:"model_relation_id" ` // 模块关联id
	CreateTime      int64  `json:"create_time" form:"create_time" `             // 收藏时间
	Model           string `json:"model" form:"model" `                         // 模块表名
	Ip              string `json:"ip" form:"ip" `                               // IP
	DeleteTime      int64  `json:"delete_time" form:"delete_time" `             // 软删除时间
	OpenId          int    `json:"open_id" form:"open_id" `                     // 商家ID

}

func (*Visit) TableName() string {
	return "visit"
}
