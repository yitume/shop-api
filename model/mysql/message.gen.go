package mysql

type Message struct {
	Id              int    `json:"id" form:"id" gorm:"primary_key"`             // 短消息索引id
	Title           string `json:"title" form:"title" `                         // 短消息标题
	Body            string `json:"body" form:"body" `                           // 消息内容,数据格式为json
	CreateTime      int64  `json:"create_time" form:"create_time" `             // 短消息发送时间
	SendState       int    `json:"send_state" form:"send_state" `               // 发送状态 0未发送 1已发送
	RelationModel   string `json:"relation_model" form:"relation_model" `       // 关联模块，并非表名
	RelationModelId int    `json:"relation_model_id" form:"relation_model_id" ` // 关联表id
	IsGroup         int    `json:"is_group" form:"is_group" `                   // 是否为群发0否1是
	TypeId          int    `json:"type_id" form:"type_id" `                     // 消息类型id
	DeleteTime      int64  `json:"delete_time" form:"delete_time" `             // 软删除时间
	OpenId          int    `json:"open_id" form:"open_id" `                     // 商家ID

}

func (*Message) TableName() string {
	return "message"
}
