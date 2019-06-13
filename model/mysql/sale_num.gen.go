package mysql

type SaleNum struct {
	Date    int64 `json:"date" form:"date" `         // 销售日期 格式：20151019
	Salenum int   `json:"salenum" form:"salenum" `   // 销量
	GoodsId int   `json:"goods_id" form:"goods_id" ` // 商品ID
	OpenId  int   `json:"open_id" form:"open_id" `   // 商家ID

}

func (*SaleNum) TableName() string {
	return "sale_num"
}
