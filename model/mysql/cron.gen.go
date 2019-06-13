package mysql

type Cron struct {
	Id      int    `json:"id" form:"id" gorm:"index"` //
	Type    int    `json:"type" form:"type" `         // 任务类型 1商品上架 2发送邮件 3优惠套装过期 4推荐展位过期
	Exeid   int    `json:"exeid" form:"exeid" `       // 关联任务的ID[如商品ID,会员ID]
	Exetime int    `json:"exetime" form:"exetime" `   // 执行时间
	Code    string `json:"code" form:"code" `         // 邮件模板CODE
	Content string `json:"content" form:"content" `   // 内容
	OpenId  int    `json:"open_id" form:"open_id" `   // 商家ID

}

func (*Cron) TableName() string {
	return "cron"
}
