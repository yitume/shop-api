package mysql

type Page struct {
	Id              int    `json:"id" form:"id" gorm:"primary_key"`           // 自增id
	Name            string `json:"name" form:"name" `                         // 模板名称
	Description     string `json:"description" form:"description" `           // 页面描述
	Body            string `json:"body" form:"body" `                         // 模板json内容
	IsPortal        int    `json:"is_portal" form:"is_portal" gorm:"index"`   // 是否为主页
	IsSystem        int    `json:"is_system" form:"is_system" `               // 系统级 默认0自定义 1系统
	BackgroundColor string `json:"background_color" form:"background_color" ` // 背景颜色
	Type            string `json:"type" form:"type" `                         // 模板类型
	CreateTime      int64  `json:"create_time" form:"create_time" `           // 创建时间
	UpdateTime      int64  `json:"update_time" form:"update_time" `           // 更新时间
	Module          string `json:"module" form:"module" `                     // 模块名字
	DeleteTime      int64  `json:"delete_time" form:"delete_time" `           // 软删除时间
	CloneFromId     int    `json:"clone_from_id" form:"clone_from_id" `       //
	OpenId          int    `json:"open_id" form:"open_id" `                   // 商家ID

}

func (*Page) TableName() string {
	return "page"
}
