package logic

import (
	"math"

	"github.com/pkg/errors"
	"github.com/thoas/go-funk"

	"github.com/yitume/shop-api/model/mysql"
)

type CartItem struct {
	Id                int                    `json:"id"`                  // cart id
	CreateTime        int64                  `json:"create_time"`         // cart create time
	GoodsSkuId        int                    `json:"goods_sku_id"`        // cart goods sku id
	GoodsNum          int                    `json:"goods_num"`           // cart goods num
	IsCheck           int                    `json:"is_check"`            // cart is check
	GoodsId           int                    `json:"goods_id"`            // goods id
	GoodsTitle        string                 `json:"goods_title"`         // goods title
	GoodsIsOnSale     int                    `json:"goods_is_on_sale"`    // goods is on sale
	GoodsFreightFee   float64                `json:"goods_freight_fee"`   // goods freight fee
	GoodsFreightId    int                    `json:"goods_freight_id"`    // goods freight id
	GoodsPayType      int                    `json:"goods_pay_type"`      // goods pay type
	GoodsPrice        float64                `json:"goods_price"`         // goods sku price
	GoodsSpec         mysql.GoodsSkuSpecJson `json:"goods_spec"`          // goods sku spec
	GoodsWeight       float64                `json:"goods_weight"`        // goods sku weight
	GoodsStock        int                    `json:"goods_stock"`         // goods sku stock
	GoodsSkuImg       string                 `json:"goods_sku_img"`       // goods sku img
	GoodsFreightAreas mysql.FreightAreasJson `json:"goods_freight_areas"` // freight areas
}

func (c CartItem) GetFreightWay() string {
	if c.GoodsFreightId == 0 {
		return "goods_freight_unified"
	}
	return "goods_freight_template"
}

func (c *CartItem) FreightFeeByAddress(address mysql.Address) (float64, error) {
	// 运费计算方式
	// 如果模板id为0，说明是freight_unified，直接用指定运费
	if c.GoodsFreightId == 0 {
		return c.GoodsFreightFee, nil
	}

	var algorithm mysql.FreightAreas
	// 如果模板id大于0，选择对应的oods_freight_template

	var flag = false
	freightAreas := c.GoodsFreightAreas
	for _, value := range freightAreas {
		if funk.ContainsInt(value.AreaIds, address.AreaId) {
			algorithm = value
			flag = true
			break
		}
	}

	if flag {
		flag = false
		for _, value := range freightAreas {
			if funk.ContainsInt(value.AreaIds, address.CityId) {
				algorithm = value
				flag = true
				break
			}
		}
	}

	// todo 要看这里逻辑
	if flag {
		flag = false
		for _, value := range freightAreas {
			if funk.ContainsInt(value.AreaIds, address.ProvinceId) {
				algorithm = value
				flag = true
				break
			}
		}
	}

	firstAmout := algorithm.FirstAmount
	firstFee := algorithm.FirstFee
	additionalAmount := algorithm.AdditionalAmount
	additionalFee := algorithm.AdditionalFee

	if c.GoodsPayType == 2 {
		weight := c.GoodsWeight * float64(c.GoodsNum)
		if weight == 0 {
			return 0, errors.New("重量不可能为0")
		}
		return firstFee + math.Ceil((weight-firstAmout)/additionalAmount)*additionalFee, nil
	}
	if c.GoodsNum == 0 {
		return 0, errors.New("个数不可能为0")
	}
	return firstFee + math.Ceil((float64(c.GoodsNum)-firstAmout)/additionalAmount)*additionalFee, nil

}
