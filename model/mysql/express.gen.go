package mysql

type Express struct {
	Id            int    `json:"id" form:"id" gorm:"primary_key"`         // 物流id
	CompanyName   string `json:"company_name" form:"company_name" `       // 公司名称
	Kuaidi100Code string `json:"kuaidi100_code" form:"kuaidi100_code" `   // 快递100Code
	TaobaoCode    string `json:"taobao_code" form:"taobao_code" `         // 淘宝100Code
	IsCommonlyUse int    `json:"is_commonly_use" form:"is_commonly_use" ` // 是否为常用物流 1 是   0 否
	CreateTime    int64  `json:"create_time" form:"create_time" `         // 添加时间
	UpdateTime    int64  `json:"update_time" form:"update_time" `         // 更新时间
	DeleteTime    int64  `json:"delete_time" form:"delete_time" `         // 软删除时间
	IsSystem      int    `json:"is_system" form:"is_system" `             // 系统级 默认0自定义 1系统
	OpenId        int    `json:"open_id" form:"open_id" `                 // 商家ID

}

func (*Express) TableName() string {
	return "express"
}
