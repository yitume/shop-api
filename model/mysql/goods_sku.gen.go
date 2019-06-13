package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type GoodsSku struct {
	Id            int              `json:"id" form:"id" gorm:"primary_key"`         // 商品id(SKU)
	GoodsId       int              `json:"goods_id" form:"goods_id" gorm:"index"`   // 商品公共表id
	Spec          GoodsSkuSpecJson `json:"spec" form:"spec" `                       // 规格json信息
	Price         float64          `json:"price" form:"price" `                     // 商品价格
	Stock         int              `json:"stock" form:"stock" `                     // 商品库存
	Code          string           `json:"code" form:"code" `                       // 商品编码
	Img           string           `json:"img" form:"img" `                         // 商品主图
	Weight        float64          `json:"weight" form:"weight" `                   // 商品重量
	Title         string           `json:"title" form:"title" `                     // 商品名称（+规格名称）
	SaleNum       int              `json:"sale_num" form:"sale_num" `               // 销售数量
	GroupSaleNum  int              `json:"group_sale_num" form:"group_sale_num" `   // 拼团销量
	SpecValueSign string           `json:"spec_value_sign" form:"spec_value_sign" ` // 规格值标识
	SpecSign      string           `json:"spec_sign" form:"spec_sign" `             // 规格标识
	CreateTime    int64            `json:"create_time" form:"create_time" `         // 商品添加时间
	UpdateTime    int64            `json:"update_time" form:"update_time" `         // 商品编辑时间
	DeleteTime    int64            `json:"delete_time" form:"delete_time" `         // 软删除时间
	OpenId        int              `json:"open_id" form:"open_id" `                 //

}

func (*GoodsSku) TableName() string {
	return "goods_sku"
}

// 请在model/mysql/addtion.go里定义GoodsSkuSpecJson结构体
func (c GoodsSkuSpecJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsSkuSpecJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
