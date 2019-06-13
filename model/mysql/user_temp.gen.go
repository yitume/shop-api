package mysql

type UserTemp struct {
	UserId           int     `json:"user_id" form:"user_id" gorm:"primary_key"`     //
	CostAveragePrice float64 `json:"cost_average_price" form:"cost_average_price" ` //
	CostTimes        int64   `json:"cost_times" form:"cost_times" `                 //
	ResentCostTime   int64   `json:"resent_cost_time" form:"resent_cost_time" `     //
	ResentVisitTime  int64   `json:"resent_visit_time" form:"resent_visit_time" `   //
	CostPrice        float64 `json:"cost_price" form:"cost_price" `                 //
	OpenId           int     `json:"open_id" form:"open_id" `                       // 商家ID

}

func (*UserTemp) TableName() string {
	return "user_temp"
}
