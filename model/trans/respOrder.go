package trans

import (
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/view"
)

type RespOrderInfo struct {
	Info              mysql.Order         `json:"info"`
	OrderCondition    view.OrderCondition `json:"order_condition"`
	OrderLog          []mysql.OrderLog    `json:"order_log"`
	ExtendOrderExtend mysql.OrderExtend   `json:"extend_order_extend"`
	ExtendOrderGoods  []FaOrderExtend     `json:"extend_order_goods"`
}
