package mysql

type Plugin struct {
	Name        string `json:"name" form:"name" `                 //
	InstallTime int64  `json:"install_time" form:"install_time" ` // 安装时间
	DbVersion   string `json:"db_version" form:"db_version" `     // 数据版本号
	OpenId      int    `json:"open_id" form:"open_id" `           // 商家ID

}

func (*Plugin) TableName() string {
	return "plugin"
}
