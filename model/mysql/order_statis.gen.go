package mysql

type OrderStatis struct {
	Id                      int64   `json:"id" form:"id" gorm:"primary_key"`                               // 统计编号(年月日)
	StartDate               int64   `json:"start_date" form:"start_date" `                                 // 开始日期
	EndDate                 int64   `json:"end_date" form:"end_date" `                                     // 结束日期
	OrderTotals             float64 `json:"order_totals" form:"order_totals" `                             // 订单金额
	OrderShippingTotals     float64 `json:"order_shipping_totals" form:"order_shipping_totals" `           // 运费
	OrderReturnTotals       float64 `json:"order_return_totals" form:"order_return_totals" `               // 退单金额
	OrderCommisTotals       float64 `json:"order_commis_totals" form:"order_commis_totals" `               // 佣金金额
	OrderCommisReturnTotals float64 `json:"order_commis_return_totals" form:"order_commis_return_totals" ` // 退还佣金
	CostTotals              float64 `json:"cost_totals" form:"cost_totals" `                               // 店铺促销活动费用
	ResultTotals            float64 `json:"result_totals" form:"result_totals" `                           // 本期应结
	CreateTime              int64   `json:"create_time" form:"create_time" `                               // 创建记录日期
	DeleteTime              int64   `json:"delete_time" form:"delete_time" `                               // 软删除时间
	OpenId                  int     `json:"open_id" form:"open_id" `                                       // 商家ID

}

func (*OrderStatis) TableName() string {
	return "order_statis"
}
