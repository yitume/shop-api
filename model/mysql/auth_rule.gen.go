package mysql

type AuthRule struct {
	Id        int    `json:"id" form:"id" gorm:"primary_key"` //
	Sign      string `json:"sign" form:"sign" `               // 规则唯一标识
	Title     string `json:"title" form:"title" `             // 规则中文名称
	Status    int    `json:"status" form:"status" `           // 状态：为1正常，为0禁用
	Type      string `json:"type" form:"type" `               //
	Condition string `json:"condition" form:"condition" `     // 规则表达式，为空表示存在就验证，不为空表示按照条件验证
	Pid       int    `json:"pid" form:"pid" `                 // 父级id
	IsSystem  int    `json:"is_system" form:"is_system" `     // 系统节点
	Sort      int    `json:"sort" form:"sort" `               // 排序
	IsDisplay int    `json:"is_display" form:"is_display" `   // 0不展示1展示
	OpenId    int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*AuthRule) TableName() string {
	return "auth_rule"
}
