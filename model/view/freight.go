package view

type Freight struct {
	GoodsSkuId int     `json:"goods_sku_id"`
	FreightFee float64 `json:"freight_fee"`
	FreightWay string  `json:"freight_way"`
}
