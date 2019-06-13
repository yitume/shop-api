package trans

type ReqAddressSetDefault struct {
	Id int `json:"id"`
}

type ReqAddressInfo struct {
	Id int `form:"id"`
}

type ReqAddressCreateInfo struct {
	Address     string `json:"address"`
	Truename    string `json:"truename"`
	MobilePhone string `json:"mobile_phone"`
	Type        string `json:"type"`
	IsDefault   int    `json:"is_default"`
	AreaId      int    `json:"area_id"`
}

type ReqAddressUpdateInfo struct {
	Id          int    `json:"id"`
	Address     string `json:"address"`
	Truename    string `json:"truename"`
	MobilePhone string `json:"mobile_phone"`
	Type        string `json:"type"`
	IsDefault   int    `json:"is_default"`
	AreaId      int    `json:"area_id"`
}

type ReqAddressDel struct {
	Id int `json:"id"`
}
