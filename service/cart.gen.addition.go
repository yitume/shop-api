package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/pkg/logic"
	"github.com/yitume/shop-api/router/mdw"
)

func (*cart) TotalNum(uid int, openId int) (cnt int) {
	type info struct {
		Total int
	}
	resp := info{}
	model.Db.Table("cart").Select("sum(goods_num) as total").Where("uid = ? AND open_id =? AND delete_time = ?", uid, openId, 0).Find(&resp)
	cnt = resp.Total
	return
}

func (*cart) ListArr(c *gin.Context, uid int, ids []int) (output []logic.CartItem) {
	condition := mysql.Conds{
		"uid": uid,
	}

	if len(ids) > 0 {
		condition["id"] = mysql.Cond{"in", ids}
	}

	cartList, _ := Cart.List(c, condition)
	var goodsSkuIds []int
	for _, value := range cartList {
		goodsSkuIds = append(goodsSkuIds, value.GoodsSkuId)
	}

	skuMap, _ := GoodsSku.ListMap(c, mysql.Conds{
		"id": mysql.Cond{"in", goodsSkuIds},
	})

	fmt.Println("skuMap", skuMap)

	var goodsIds []int
	for _, value := range skuMap {
		goodsIds = append(goodsIds, value.GoodsId)
	}

	goodsMap, _ := Goods.ListMap(c, mysql.Conds{
		"id": mysql.Cond{"in", goodsIds},
	})

	var freightIds []int
	for _, value := range goodsMap {
		freightIds = append(freightIds, value.FreightId)
	}
	freightMap, _ := Freight.ListMap(c, mysql.Conds{
		"id": mysql.Cond{"in", freightIds},
	})

	output = make([]logic.CartItem, 0)
	for _, cartInfo := range cartList {
		tmp := logic.CartItem{
			Id:         cartInfo.Id,
			CreateTime: cartInfo.CreateTime,
			GoodsSkuId: cartInfo.GoodsSkuId,
			GoodsNum:   cartInfo.GoodsNum,
			IsCheck:    cartInfo.IsCheck,
		}
		skuInfo := skuMap[cartInfo.GoodsSkuId]
		goodsInfo := goodsMap[skuInfo.GoodsId]
		freightInfo := freightMap[goodsInfo.FreightId]

		tmp.GoodsId = skuInfo.GoodsId
		tmp.GoodsTitle = goodsInfo.Title
		tmp.GoodsIsOnSale = goodsInfo.IsOnSale
		tmp.GoodsFreightFee = goodsInfo.FreightFee
		tmp.GoodsFreightId = goodsInfo.FreightId
		tmp.GoodsPayType = goodsInfo.PayType
		tmp.GoodsPrice = goodsInfo.Price
		tmp.GoodsSpec = skuInfo.Spec
		tmp.GoodsWeight = skuInfo.Weight
		tmp.GoodsStock = skuInfo.Stock
		tmp.GoodsSkuImg = skuInfo.Img
		tmp.GoodsFreightAreas = freightInfo.Areas

		output = append(output, tmp)
	}
	return
}

func (*cart) InfoA(c *gin.Context, uid int, goodsSkuId int) (output logic.CartItem, err error) {
	cartInfo, err := Cart.InfoX(c, mysql.Conds{
		"uid":          uid,
		"goods_sku_id": goodsSkuId,
		"open_id":      mdw.OpenId(c),
	})

	if err != nil {
		return
	}

	skuInfo, err := GoodsSku.InfoX(c, mysql.Conds{
		"id": goodsSkuId,
	})
	if err != nil {
		return
	}
	goodsInfo, err := Goods.InfoX(c, mysql.Conds{
		"id": skuInfo.GoodsId,
	})
	if err != nil {
		return
	}
	freightInfo, err := Freight.InfoX(c, mysql.Conds{
		"id": goodsInfo.FreightId,
	})
	if err != nil {
		return
	}

	output = logic.CartItem{
		Id:                cartInfo.Id,
		CreateTime:        cartInfo.CreateTime,
		GoodsSkuId:        cartInfo.GoodsSkuId,
		GoodsNum:          cartInfo.GoodsNum,
		IsCheck:           cartInfo.IsCheck,
		GoodsId:           skuInfo.GoodsId,
		GoodsTitle:        goodsInfo.Title,
		GoodsIsOnSale:     goodsInfo.IsOnSale,
		GoodsFreightFee:   goodsInfo.FreightFee,
		GoodsFreightId:    goodsInfo.FreightId,
		GoodsPayType:      goodsInfo.PayType,
		GoodsPrice:        goodsInfo.Price,
		GoodsSpec:         skuInfo.Spec,
		GoodsWeight:       skuInfo.Weight,
		GoodsStock:        skuInfo.Stock,
		GoodsSkuImg:       skuInfo.Img,
		GoodsFreightAreas: freightInfo.Areas,
	}

	return
}
