package mysql

type MessageState struct {
	Id           int   `json:"id" form:"id" gorm:"primary_key"`       //
	ToUserId     int   `json:"to_user_id" form:"to_user_id" `         // 接收者id
	MessageId    int   `json:"message_id" form:"message_id" `         // 信息id
	ReadState    int   `json:"read_state" form:"read_state" `         // 0 未读 1已读
	DelState     int   `json:"del_state" form:"del_state" `           // 0未删 1已删
	ReadTime     int64 `json:"read_time" form:"read_time" `           // 读取时间
	DelTime      int64 `json:"del_time" form:"del_time" `             // 删除时间
	AppPushState int   `json:"app_push_state" form:"app_push_state" ` // app是否推送0未1是
	CreateTime   int64 `json:"create_time" form:"create_time" `       // 创建时间
	DeleteTime   int64 `json:"delete_time" form:"delete_time" `       // 软删除时间
	OpenId       int   `json:"open_id" form:"open_id" `               // 商家ID

}

func (*MessageState) TableName() string {
	return "message_state"
}
