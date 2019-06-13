package mysql

type OrderLog struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // 主键
	OrderId    int    `json:"order_id" form:"order_id" `       // 订单id
	Msg        string `json:"msg" form:"msg" `                 // 文字描述
	CreateTime int64  `json:"create_time" form:"create_time" ` // 处理时间
	Role       string `json:"role" form:"role" `               // 操作角色
	User       string `json:"user" form:"user" `               // 操作人
	OrderState int    `json:"order_state" form:"order_state" ` // 订单状态：0(已取消)10:未付款;20:已付款;30:已发货;40:已收货;
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商家ID

}

func (*OrderLog) TableName() string {
	return "order_log"
}
