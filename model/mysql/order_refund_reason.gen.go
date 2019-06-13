package mysql

type OrderRefundReason struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // id
	Title      string `json:"title" form:"title" `             // 标题
	Type       int    `json:"type" form:"type" `               // 1未收到货 2已收到货
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*OrderRefundReason) TableName() string {
	return "order_refund_reason"
}
