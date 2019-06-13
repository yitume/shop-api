package mysql

type MessageType struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` //
	Title      string `json:"title" form:"title" `             // 通知类型标题
	Remark     string `json:"remark" form:"remark" `           // 通知类型的备注
	Status     int    `json:"status" form:"status" `           // 状态0关闭，1开启
	Model      string `json:"model" form:"model" `             // 关联模型
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*MessageType) TableName() string {
	return "message_type"
}
