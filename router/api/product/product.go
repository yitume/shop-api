package product

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/router/base"
)

// List 商品列表
func List(c *gin.Context) {
	var products []mysql.Product
	pageNo, err := strconv.Atoi(c.Query("pageNo"))
	if err != nil || pageNo < 1 {
		pageNo = 1
	}

	offset := (pageNo - 1) * 20

	//默认按创建时间，降序来排序
	var orderStr = "created_at"
	if c.Query("order") == "1" {
		orderStr = "total_sale"
	} else if c.Query("order") == "2" {
		orderStr = "created_at"
	}
	if c.Query("asc") == "1" {
		orderStr += " asc"
	} else {
		orderStr += " desc"
	}

	cateID, err := strconv.Atoi(c.Query("cateId"))

	if err != nil {
		base.JSON(c, base.MsgErr, "分类ID不正确")
		return
	}

	var category mysql.Category

	if model.Db.First(&category, cateID).Error != nil {
		base.JSON(c, base.MsgErr, "分类ID不正确")
		return
	}

	pageSize := 20
	queryErr := model.Db.Offset(offset).Limit(pageSize).Order(orderStr).Find(&products).Error

	if queryErr != nil {
		base.JSON(c, base.MsgErr, "err")
		return
	}

	for i := 0; i < len(products); i++ {
		err := model.Db.First(&products[i].Image, products[i].ImageID).Error
		if err != nil {
			base.JSON(c, base.MsgErr, "err")
			return
		}
	}

	base.JSON(c, base.MsgOk, "ok", gin.H{
		"products": products,
	})
}

// Info 获取商品信息
func Info(c *gin.Context) {
	reqStartTime := time.Now()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.JSON(c, base.MsgErr, "错误的商品id")
		return
	}

	var product mysql.Product

	if model.Db.First(&product, id).Error != nil {
		base.JSON(c, base.MsgErr, "错误的商品id")
		return
	}

	if model.Db.First(&product.Image, product.ImageID).Error != nil {
		product.Image = mysql.Image{}
	}

	var imagesSQL []uint
	if err := json.Unmarshal([]byte(product.ImageIDs), &imagesSQL); err == nil {
		var images []mysql.Image
		if model.Db.Where("id in (?)", imagesSQL).Find(&images).Error != nil {
			product.Images = nil
		} else {
			product.Images = images
		}
	} else {
		product.Images = nil
	}

	if err := model.Db.Model(&product).Related(&product.Categories, "categories").Error; err != nil {
		base.JSON(c, base.MsgErr, "err")
		return
	}

	if product.HasProperty {
		if err := model.Db.Model(&product).Related(&product.Properties).Error; err != nil {
			base.JSON(c, base.MsgErr, "err")
			return
		}

		for i := 0; i < len(product.Properties); i++ {
			property := product.Properties[i]
			if err := model.Db.Model(&property).Related(&property.PropertyValues).Error; err != nil {
				base.JSON(c, base.MsgErr, "err")

				return
			}
			product.Properties[i] = property
		}

		if err := model.Db.Model(&product).Related(&product.Inventories).Error; err != nil {
			base.JSON(c, base.MsgErr, "err")

			return
		}

		for i := 0; i < len(product.Inventories); i++ {
			inventory := product.Inventories[i]
			if err := model.Db.Model(&inventory).Related(&inventory.PropertyValues, "property_values").Error; err != nil {
				base.JSON(c, base.MsgErr, "err")
				return
			}
			product.Inventories[i] = inventory
		}
	}

	fmt.Println("duration: ", time.Now().Sub(reqStartTime).Seconds())
	base.JSON(c, base.MsgOk, gin.H{
		"product": product,
	})

}
