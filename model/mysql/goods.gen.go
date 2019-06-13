package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type Goods struct {
	Id                 int                      `json:"id" form:"id" gorm:"primary_key"`                   // 商品公共表id
	Title              string                   `json:"title" form:"title" `                               // 商品名称
	Images             GoodsImagesJson          `json:"images" form:"images" `                             // 商品图片 默认第一个为封面图片
	CategoryIds        GoodsCategoryIdsJson     `json:"category_ids" form:"category_ids" `                 // 商品分类
	BaseSaleNum        int                      `json:"base_sale_num" form:"base_sale_num" `               // 基础销量
	Body               GoodsBodyJson            `json:"body" form:"body" `                                 // 商品内容
	IsOnSale           int                      `json:"is_on_sale" form:"is_on_sale" `                     // 是否需上架出售 0 否 1 是
	ImageSpecId        int                      `json:"image_spec_id" form:"image_spec_id" `               // 使用图片的规格id
	ImageSpecImages    GoodsImageSpecImagesJson `json:"image_spec_images" form:"image_spec_images" `       // 规格图片集合，废弃
	SkuList            GoodsSkuListJson         `json:"sku_list" form:"sku_list" `                         // sku商品集合，数组
	CreateTime         int64                    `json:"create_time" form:"create_time" `                   // 创建时间
	Price              float64                  `json:"price" form:"price" `                               // 商品价格
	UpdateTime         int64                    `json:"update_time" form:"update_time" `                   // 修改时间
	EvaluationGoodStar int                      `json:"evaluation_good_star" form:"evaluation_good_star" ` // 好评星级
	EvaluationCount    int                      `json:"evaluation_count" form:"evaluation_count" `         // 评价数
	Stock              int                      `json:"stock" form:"stock" `                               // goods表库存之和
	SaleNum            int                      `json:"sale_num" form:"sale_num" `                         // 销售量
	GroupSaleNum       int                      `json:"group_sale_num" form:"group_sale_num" `             // 拼团销量
	SaleTime           int64                    `json:"sale_time" form:"sale_time" `                       // 开售时间
	DeleteTime         int64                    `json:"delete_time" form:"delete_time" `                   // 软删除时间
	SpecList           GoodsSpecListJson        `json:"spec_list" form:"spec_list" `                       // {&quot;type&quot;:&quot;slice&quot;,&quot;data&quot;:&quot;spec&quot;}
	Img                string                   `json:"img" form:"img" `                                   // 封面图
	PayType            int                      `json:"pay_type" form:"pay_type" `                         // 计算方式：1 按件数 2 按重量
	FreightFee         float64                  `json:"freight_fee" form:"freight_fee" `                   // 运费
	FreightId          int                      `json:"freight_id" form:"freight_id" `                     // 运费模板id
	OpenId             int                      `json:"open_id" form:"open_id" `                           // 商户id

}

func (*Goods) TableName() string {
	return "goods"
}

// 请在model/mysql/addtion.go里定义GoodsImagesJson结构体
func (c GoodsImagesJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsImagesJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义GoodsCategoryIdsJson结构体
func (c GoodsCategoryIdsJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsCategoryIdsJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义GoodsBodyJson结构体
func (c GoodsBodyJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsBodyJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义GoodsImageSpecImagesJson结构体
func (c GoodsImageSpecImagesJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsImageSpecImagesJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义GoodsSkuListJson结构体
func (c GoodsSkuListJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsSkuListJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义GoodsSpecListJson结构体
func (c GoodsSpecListJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsSpecListJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
