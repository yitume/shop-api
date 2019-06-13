package trans

import "encoding/json"

type Page struct {
	Id              int             `json:"id"`               // 自增id
	Name            string          `json:"name"`             // 模板名称
	Description     string          `json:"description"`      // 页面描述
	Body            json.RawMessage `json:"body"`             // 模板json内容
	IsPortal        int             `json:"is_portal"`        // 是否为主页
	IsSystem        int             `json:"is_system"`        // 系统级 默认0自定义 1系统
	BackgroundColor string          `json:"background_color"` // 背景颜色
	Type            string          `json:"type"`             // 模板类型
	CreateTime      int64           `json:"create_time"`      // 创建时间
	UpdateTime      int64           `json:"update_time"`      // 更新时间
	Module          string          `json:"module"`           // 模块名字
	CloneFromId     int             `json:"clone_from_id"`    //
}
