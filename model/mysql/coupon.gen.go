package mysql

type Coupon struct {
	Id              int     `json:"id" form:"id" gorm:"primary_key"`             // id
	Title           string  `json:"title" form:"title" `                         // 名称
	StartTime       int64   `json:"start_time" form:"start_time" `               // 开始时间
	EndTime         int64   `json:"end_time" form:"end_time" `                   // 结束时间
	Denomination    float64 `json:"denomination" form:"denomination" `           // 面额
	TimeType        int64   `json:"time_type" form:"time_type" `                 // 有效时间类型 默认0 XXX天内有效 1固定时间段
	EffectiveDays   int     `json:"effective_days" form:"effective_days" `       // XXX天内有效
	UseStartTime    int64   `json:"use_start_time" form:"use_start_time" `       // 使用开始时间
	UseEndTime      int64   `json:"use_end_time" form:"use_end_time" `           // 使用结束时间
	Number          int     `json:"number" form:"number" `                       // 发放数量
	LimitType       int     `json:"limit_type" form:"limit_type" `               // 使用条件 默认0不限制 1满XXX使用
	LimitPrice      float64 `json:"limit_price" form:"limit_price" `             // 满XXX使用
	ReceiveLimitNum int     `json:"receive_limit_num" form:"receive_limit_num" ` // 每人限领 0不限制
	Level           int     `json:"level" form:"level" `                         // 级别 默认0全店 1商品级
	Partake         string  `json:"partake" form:"partake" `                     // json 优惠参与 折扣discount 满减fullcut 优惠券coupon 废弃
	CreateTime      int64   `json:"create_time" form:"create_time" `             // 创建时间
	DeleteTime      int64   `json:"delete_time" form:"delete_time" `             // 软删除时间
	OpenId          int     `json:"open_id" form:"open_id" `                     // 商户id

}

func (*Coupon) TableName() string {
	return "coupon"
}
