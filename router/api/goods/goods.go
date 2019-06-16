package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/pkg/common/code"
	"time"

	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/yitume/shop-api/service"
)

func GoodsInfo(c *gin.Context) {
	req := trans.ReqGoodsInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgGoodsInfoReqErr)
		return
	}

	goodsInfo, err := service.Goods.InfoX(c, mysql.Conds{
		"id":      req.Id,
		"open_id": mdw.OpenId(c),
	})
	if err != nil {
		base.JSON(c, code.MsgGoodsInfoMysqlErr)
		return
	}

	skuList, _ := service.GoodsSku.List(c, mysql.Conds{
		"goods_id": goodsInfo.Id,
		"open_id":  mdw.OpenId(c),
	}, "id desc")

	resp3 := make([]trans.FaGoodsSku, 0)
	for _, skuInfo := range skuList {
		resp3 = append(resp3, trans.FaGoodsSku{
			Id:            skuInfo.Id,
			GoodsId:       skuInfo.GoodsId,
			Spec:          skuInfo.Spec,
			Price:         skuInfo.Price,
			Stock:         skuInfo.Stock,
			Code:          skuInfo.Code,
			Img:           skuInfo.Img,
			Weight:        skuInfo.Weight,
			Title:         skuInfo.Title,
			SaleNum:       skuInfo.SaleNum,
			GroupSaleNum:  skuInfo.GroupSaleNum,
			SpecSign:      skuInfo.SpecSign,
			SpecValueSign: skuInfo.SpecValueSign,
		})
	}

	resp := trans.FaGoods{
		Id:                 goodsInfo.Id,
		Title:              goodsInfo.Title,
		Images:             goodsInfo.Images,
		CategoryIds:        goodsInfo.CategoryIds,
		BaseSaleNum:        goodsInfo.BaseSaleNum,
		Body:               goodsInfo.Body,
		IsOnSale:           goodsInfo.IsOnSale,
		ImageSpecId:        goodsInfo.ImageSpecId,
		ImageSpecImages:    goodsInfo.ImageSpecImages,
		SkuList:            resp3,
		CreateTime:         goodsInfo.CreateTime,
		Price:              goodsInfo.Price,
		UpdateTime:         goodsInfo.UpdateTime,
		EvaluationCount:    goodsInfo.EvaluationCount,
		EvaluationGoodStar: goodsInfo.EvaluationGoodStar,
		Stock:              goodsInfo.Stock,
		SaleNum:            goodsInfo.SaleNum,

		GroupSaleNum: goodsInfo.GroupSaleNum,
		SaleTime:     goodsInfo.SaleTime,

		SpecList: goodsInfo.SpecList,
		Img:      goodsInfo.Img,
		PayType:  goodsInfo.PayType,

		FreightFee: goodsInfo.FreightFee,
		FreightId:  goodsInfo.FreightId,
	}
	base.JSON(c, code.MsgOk, map[string]interface{}{
		"info": resp,
	})

}

func GoodsList(c *gin.Context) {
	req := trans.ReqGoodsListCustom{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgGoodsListReqErr)
		return
	}

	output := make([]mysql.Goods, 0)
	model.Db.Table("goods as a").Select("*").Joins("LEFT JOIN goods_category_ids as b ON a.id = b.goods_id").Where("b.category_id in (?) and a.sale_time < ? and a.is_on_sale = ?", req.CategoryIds, time.Now().Unix(), 1).Offset((req.Page - 1) * req.Rows).Limit(req.Rows).Find(&output)

	base.JSONList(c, output, 0)
}
