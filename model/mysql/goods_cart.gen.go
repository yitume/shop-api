package mysql

type GoodsCart struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"`       //
	Uniacid    int    `json:"uniacid" form:"uniacid" `               //
	Couponid   string `json:"couponid" form:"couponid" gorm:"index"` //
	Storeid    int    `json:"storeid" form:"storeid" `               //
	DeleteTime int64  `json:"delete_time" form:"delete_time" `       // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `               // 商户id

}

func (*GoodsCart) TableName() string {
	return "goods_cart"
}
