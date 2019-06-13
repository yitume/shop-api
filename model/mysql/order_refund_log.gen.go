package mysql

type OrderRefundLog struct {
	Id               int    `json:"id" form:"id" gorm:"primary_key"`               // 主键
	OrderRefundId    int    `json:"order_refund_id" form:"order_refund_id" `       // 订单id
	UserId           int    `json:"user_id" form:"user_id" `                       // 操作人
	Role             string `json:"role" form:"role" `                             // 操作角色
	OrderRefundState int    `json:"order_refund_state" form:"order_refund_state" ` // 平台处理状态 默认0处理中(未处理) 10拒绝(驳回) 20同意 30成功(已完成) 50取消(用户主动撤销) 51取消(用户主动收货)
	Msg              string `json:"msg" form:"msg" `                               // 文字描述
	CreateTime       int64  `json:"create_time" form:"create_time" `               // 操作时间
	DeleteTime       int64  `json:"delete_time" form:"delete_time" `               // 软删除时间
	OpenId           int    `json:"open_id" form:"open_id" `                       // 商家ID

}

func (*OrderRefundLog) TableName() string {
	return "order_refund_log"
}
