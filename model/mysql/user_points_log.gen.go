package mysql

type UserPointsLog struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // 主键
	Msg        string `json:"msg" form:"msg" `                 // 文字描述
	CreateTime int64  `json:"create_time" form:"create_time" ` // 处理时间
	UserId     int    `json:"user_id" form:"user_id" `         //
	Points     int    `json:"points" form:"points" `           // 积分变化的数量
	Type       string `json:"type" form:"type" `               // 类型名字
	Username   string `json:"username" form:"username" `       // 会员名称
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*UserPointsLog) TableName() string {
	return "user_points_log"
}
