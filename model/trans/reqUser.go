package trans

type ReqUserAddressList struct {
	Token string `form:"token"  binding:"required"`
}

type ReqUserAddressDefault struct {
	Token string `form:"token"  binding:"required"`
}

type ReqUserAddressCreate struct {
	Token      string `json:"token"  binding:"required"`
	ProvinceId uint32 `json:"provinceId"  binding:"required"`
	CityId     uint32 `json:"cityId"  binding:"required"`
	DistrictId uint32 `json:"districtId"`
	LinkMan    string `json:"linkMan"`
	Address    string `json:"address"`
	Mobile     string `json:"mobile"`
	Code       string `json:"code"`
	IsDefault  uint32 `json:"isDefault"`
}

type ReqUserLogin struct {
	WechatMiniCode string `json:"wechat_mini_code"`
}

type ReqUserReigster struct {
	RegisterType    string `json:"register_type"`
	WechatMiniParam struct {
		Code          string `json:"code"`
		RawData       string `json:"raw_data"`
		EncryptedData string `json:"encrypted_data"`
		Signature     string `json:"signature"`
		Iv            string `json:"iv"`
	} `json:"wechat_mini_param"`
}
