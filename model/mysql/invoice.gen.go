package mysql

type Invoice struct {
	Id                          int    `json:"id" form:"id" gorm:"primary_key"`                                       // 索引id
	UserId                      int    `json:"user_id" form:"user_id" `                                               // 会员ID
	Type                        string `json:"type" form:"type" `                                                     // 1普通发票2增值税发票
	Title                       string `json:"title" form:"title" `                                                   // 发票抬头[普通发票]
	Content                     string `json:"content" form:"content" `                                               // 发票内容[普通发票]
	CompanyName                 string `json:"company_name" form:"company_name" `                                     // 单位名称
	CompanyCode                 string `json:"company_code" form:"company_code" `                                     // 纳税人识别号
	CompanyRegisterAddress      string `json:"company_register_address" form:"company_register_address" `             // 注册地址
	CompanyRegisterPhone        string `json:"company_register_phone" form:"company_register_phone" `                 // 注册电话
	CompanyRegisterBrankName    string `json:"company_register_brank_name" form:"company_register_brank_name" `       // 开户银行
	CompanyRegisterBrankAccount string `json:"company_register_brank_account" form:"company_register_brank_account" ` // 银行帐户
	ReceiveName                 string `json:"receive_name" form:"receive_name" `                                     // 收票人姓名
	ReceivePhone                string `json:"receive_phone" form:"receive_phone" `                                   // 收票人手机号
	ReceiveProvince             string `json:"receive_province" form:"receive_province" `                             // 收票人省份
	ReceiveAddress              string `json:"receive_address" form:"receive_address" `                               // 送票地址
	ConsumptionType             string `json:"consumption_type" form:"consumption_type" `                             // 发票消费类型
	IsDefault                   int    `json:"is_default" form:"is_default" `                                         // 默认发票
	DeleteTime                  int64  `json:"delete_time" form:"delete_time" `                                       // 软删除时间
	OpenId                      int    `json:"open_id" form:"open_id" `                                               // 商家ID

}

func (*Invoice) TableName() string {
	return "invoice"
}
