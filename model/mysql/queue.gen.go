package mysql

type Queue struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` //
	Name       string `json:"name" form:"name" `               // 队列名字(备注名)
	Raw        string `json:"raw" form:"raw" `                 // 生数据
	CreateTime int64  `json:"create_time" form:"create_time" ` // 创建时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Queue) TableName() string {
	return "queue"
}
