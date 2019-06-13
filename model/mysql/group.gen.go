package mysql

type Group struct {
	Id             int    `json:"id" form:"id" gorm:"primary_key"`           // id
	Title          string `json:"title" form:"title" `                       // 拼团活动名称
	Sn             string `json:"sn" form:"sn" `                             // 活动编号  预留
	TimeOver       string `json:"time_over" form:"time_over" `               // 拼团时限 格式: 小时:分钟
	TimeOverDay    int64  `json:"time_over_day" form:"time_over_day" `       // 拼团时限-天
	TimeOverHour   int64  `json:"time_over_hour" form:"time_over_hour" `     // 拼团时限-小时
	TimeOverMinute int64  `json:"time_over_minute" form:"time_over_minute" ` // 拼团时限-分钟
	StartTime      int64  `json:"start_time" form:"start_time" `             // 活动开始时间
	EndTime        int64  `json:"end_time" form:"end_time" `                 // 活动结束时间
	LimitBuyNum    int    `json:"limit_buy_num" form:"limit_buy_num" `       // 拼团人数
	LimitGroupNum  int    `json:"limit_group_num" form:"limit_group_num" `   // 每位用户可进行的团数
	LimitGoodsNum  int    `json:"limit_goods_num" form:"limit_goods_num" `   // 用户每次参团时可购买件数
	GoodsId        int    `json:"goods_id" form:"goods_id" `                 // 商品id
	IsShow         int    `json:"is_show" form:"is_show" `                   // 是否正在执行 0未执行 1执行
	OverType       int    `json:"over_type" form:"over_type" `               // 结束时间格式  0长期 1定期(一年) 预留
	CreateTime     int64  `json:"create_time" form:"create_time" `           // 创建时间
	UpdateTime     int64  `json:"update_time" form:"update_time" `           // 更新时间
	DeleteTime     int64  `json:"delete_time" form:"delete_time" `           // 软删除时间
	OpenId         int    `json:"open_id" form:"open_id" `                   // 商家ID

}

func (*Group) TableName() string {
	return "group"
}
