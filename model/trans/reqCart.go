package trans

type ReqCartInfo struct {
	GoodsSkuId int `form:"goods_sku_id"`
}

type ReqCartCreateInfo struct {
	GoodsSkuId int `json:"goods_sku_id"`
	Quantity   int `json:"quantity"`
}

type ReqCartUpdateInfo struct {
	GoodsSkuId int `json:"goods_sku_id"`
	Quantity   int `json:"quantity"`
}

type ReqCartDel struct {
	GoodsSkuIds []int `json:"goods_sku_ids"`
}

type ReqCartCheck struct {
	GoodsSkuIds []int `json:"goods_sku_ids"`
	IsCheck     int   `json:"is_check"`
}
