package mysql

type GoodBaseInfo struct {
	Id          uint32  `json:"id"`
	Name        string  `json:"name"`
	Pic         string  `json:"pic"`
	Cid         uint32  `json:"cid"`
	LogisticsId uint32  `json:"logisticsId"` // 快递id
	Stores      uint32  `json:"stores"`      // 剩余货量
	Views       uint32  `json:"views"`       // 查看次数
	MinPrice    float64 `json:"minPrice"`
	MinScore    uint32  `json:"minScore"`
	Content     string  `json:"content"`

	NumOrders         uint32 `json:"numOrders"`         //订单数
	NumGoodReputation uint32 `json:"numGoodReputation"` //好评数

	CommissionType uint32 `json:"commissionType"` // 任务类型 1 分享积分奖励， 2 分享现金奖励
	Commission     uint32 `json:"commission"`     // 奖励数字

	Slide string `json:"-"`

	IsGroupon    uint32  `json:"isGroupon"`
	GrouponPrice float64 `json:"grouponPrice"`
}

func (GoodBaseInfo) TableName() string {
	return "goods"
}
