package trans

type ReqBuyCalculate struct {
	AddressId int   `json:"address_id"`
	CartIds   []int `json:"cart_ids"`
}

type ReqBuyCreate struct {
	CartIds   []int  `json:"cart_ids"`
	AddressId int    `json:"address_id"`
	Message   string `json:"message"`
}

type ReqBuyPay struct {
	OrderType      string `json:"order_type"`
	PaySn          string `json:"pay_sn"`
	PaymentCode    string `json:"payment_code"`
	PaymentChannel string `json:"payment_channel"`
	OpenID         string `json:"open_id"`
}
