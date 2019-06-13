package trans

type ReqOrderInfo struct {
	Id int `form:"id"`
}

type ReqOrderCancel struct {
	Id          int `json:"id"`
	StateRemark int `json:"state_remark"`
}

type ReqOrderListInfo struct {
	Page      int    `form:"page"`
	Rows      int    `form:"rows"`
	StateType string `form:"state_type"`
}

type ReqOrderConfirmReceipt struct {
	Id          int    `json:"id"`
	StateRemark string `json:"state_remark"`
}

type ReqOrderGoodsInfo struct {
	Id int `json:"id"`
}

type ReqOrderLogistics struct {
	Id int `json:"id"`
}
