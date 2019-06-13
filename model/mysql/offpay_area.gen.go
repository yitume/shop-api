package mysql

type OffpayArea struct {
	Id     int    `json:"id" form:"id" gorm:"index"` // id
	AreaId string `json:"area_id" form:"area_id" `   // 县ID组合
	OpenId int    `json:"open_id" form:"open_id" `   // 商家ID

}

func (*OffpayArea) TableName() string {
	return "offpay_area"
}
