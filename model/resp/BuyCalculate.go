package resp

import "github.com/yitume/shop-api/model/view"

type BuyCalculate struct {
	GoodsAmount        float64        `json:"goods_amount"`
	PayAmount          float64        `json:"pay_amount"`
	GoodsFreightList   []view.Freight `json:"goods_freight_list"`
	FreightUnifiedFee  float64        `json:"freight_unified_fee"`
	FreightTemplateFee float64        `json:"freight_template_fee"`
	PayFreightFee      float64        `json:"pay_freight_fee"`
	SubTotal           float64        `json:"sub_total"`
	GoodsNum           int            `json:"-"`
}
