package shop

import (
	"github.com/yitume/shop-api/model/resp"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/yitume/shop-api/service"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Info(c *gin.Context) {
	info, _ := service.Shop.InfoX(c, mysql.Conds{
		"open_id": mdw.OpenId(c),
	})

	base.JSON(c, base.MsgOk, resp.ShopInfo{
		Name:                         info.Name,
		Logo:                         info.Logo,
		ContactNumber:                info.ContactNumber,
		Description:                  info.Description,
		ColorScheme:                  info.ColorScheme,
		PortalTemplateId:             info.PortalTemplateId,
		WechatPlatformQr:             info.WechatPlatformQr,
		GoodsCategoryStyle:           info.GoodsCategoryStyle,
		Host:                         info.Host,
		OrderAutoCloseExpires:        info.OrderAutoCloseExpires,
		OrderAutoConfirmExpires:      info.OrderAutoConfirmExpires,
		OrderAutoCloseRefoundExpires: info.OrderAutoCloseRefoundExpires,
	})
}
