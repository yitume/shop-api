package mysql

type Version struct {
	Id          int    `json:"id" form:"id" gorm:"primary_key"`   //
	Version     string `json:"version" form:"version" `           // 版本号
	UpdateState string `json:"update_state" form:"update_state" ` // required必须更新optional可选noneed不需要更新
	DownloadUrl string `json:"download_url" form:"download_url" ` // 下载地址
	Platform    string `json:"platform" form:"platform" `         // ios/android
	CreateTime  int64  `json:"create_time" form:"create_time" `   // 创建时间
	PublishTime int64  `json:"publish_time" form:"publish_time" ` // 发布时间
	Description string `json:"description" form:"description" `   // 更新说明
	DeleteTime  int64  `json:"delete_time" form:"delete_time" `   // 软删除时间
	OpenId      int    `json:"open_id" form:"open_id" `           // 商家ID

}

func (*Version) TableName() string {
	return "version"
}
