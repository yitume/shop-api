package category

import (
	"github.com/gin-gonic/gin"
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/router/base"
)

// List 公开的分类列表
func List(c *gin.Context) {
	var categories []mysql.Category

	if err := model.Db.Where("status = 1").Order("sequence asc").Find(&categories).Error; err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	base.JSON(c, base.MsgOk, gin.H{
		"categories": categories,
	})
}
