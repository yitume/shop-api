package mysql

type Upload struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` //
	FileName   string `json:"file_name" form:"file_name" `     // 文件名
	FileSize   int    `json:"file_size" form:"file_size" `     // 文件大小
	FileType   string `json:"file_type" form:"file_type" `     // 文件类型
	FilePath   string `json:"file_path" form:"file_path" `     // 文件存放路径
	UserId     int    `json:"user_id" form:"user_id" `         // 用户id
	CreateTime int64  `json:"create_time" form:"create_time" ` // 创建时间
	Type       string `json:"type" form:"type" `               // 标识
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Upload) TableName() string {
	return "upload"
}
