package trans

type ReqShopGoodList struct {
	Cid      int    `form:"cid"`
	Name     string `form:"name"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
}

type ReqShopGoodInfo struct {
	Id int `form:"id" binding:"required"`
}

// token:2af1a72e-2fe9-40f0-8ad7-b26d7a2473ed
// goodsJson:[{"goodsId":100804,"number":1,"propertyChildIds":"","logisticsType":0, "inviter_id":0}]
// remark:
// calculate:true
type ReqShopOrder struct {
	Token     string                  `json:"token" binding:"required"`
	GoodsJson []ReqShopOrderGoodsJson `json:"goodsJson"  binding:"required"`
	Remark    string                  `json:"remark"`
}

type ReqShopOrderGoodsJson struct {
	GoodsId          uint32 `json:"goodsId"`
	Number           uint32 `json:"number"`
	PropertyChildIds string `json:"propertyChildIds"`
	LogisticsType    uint32 `json:"logisticsType"`
	InviterId        uint32 `json:"inviterId"`
}

type ReqShopOrderList struct {
	Token  string `form:"token" binding:"required"`
	Status int    `form:"status" binding:"required"`
}

type ReqShopOrderCancel struct {
	Token   string `json:"token" binding:"required"`
	OrderNo string `json:"orderNo" binding:"required"`
}
