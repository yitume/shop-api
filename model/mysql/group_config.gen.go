package mysql

type GroupConfig struct {
	Refunmoney int `json:"refunmoney" form:"refunmoney" ` // 退款方式: 1,自动 2,手动
	OpenId     int `json:"open_id" form:"open_id" `       // 商家ID

}

func (*GroupConfig) TableName() string {
	return "group_config"
}
