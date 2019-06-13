package trans

type ReqGoodsOnSale struct {
	Ids []int `json:"ids"`
}

type ReqGoodsInfo struct {
	Id int `form:"id"`
}

type ReqGoodsListCustom struct {
	CategoryIds []int `json:"category_ids"`
	Page        int   `json:"page"`
	Rows        int   `json:"rows"`
}

type ReqGoodscategoryCreate struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
	Pid  int    `json:"pid"`
}

type ReqGoodscategoryInfo struct {
	Id int `form:"id"`
}

type ReqGoodscategoryUpdate struct {
	Id   int    `json:"id"`
	Icon string `json:"icon"`
	Name string `json:"name"`
	Pid  int    `json:"pid"`
}

type ReqGoodsspecCreate struct {
	Name string `json:"name"`
}

type ReqGoodsspecvalueCreate struct {
	SpecId int    `json:"spec_id"`
	Name   string `json:"name"`
}
