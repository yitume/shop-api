package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/pkg/bootstrap"
	"github.com/yitume/shop-api/router/api/address"
	"github.com/yitume/shop-api/router/api/area"
	"github.com/yitume/shop-api/router/api/buy"
	"github.com/yitume/shop-api/router/api/cart"
	"github.com/yitume/shop-api/router/api/collect"
	"github.com/yitume/shop-api/router/api/goods"
	"github.com/yitume/shop-api/router/api/order"
	"github.com/yitume/shop-api/router/api/page"
	"github.com/yitume/shop-api/router/api/shop"
	"github.com/yitume/shop-api/router/api/user"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/yitume/shop-api/router/mdw/ginzap"
	"github.com/yitume/shop-api/router/mdw/wechat"
)

func InitApi() *gin.Engine {
	gin.SetMode(bootstrap.Conf.App.Mode)

	r := gin.New()
	r.Use(ginzap.Ginzap(model.Logger.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(model.Logger.Logger, true))

	store := persistence.NewInMemoryStore(time.Second)
	domainGroup := r.Group("/api/:domain", mdw.Domain())
	{
		pageGrp := domainGroup.Group("/page")
		{
			pageGrp.GET("/portal", page.Portal)
			pageGrp.GET("/info")
		}

		goodsGrp := domainGroup.Group("/goods")
		{
			goodsGrp.POST("/list", goods.GoodsList)
			goodsGrp.GET("/info", goods.GoodsInfo)
		}

		goodscategoryGrp := domainGroup.Group("/goodscategory")
		{
			goodscategoryGrp.GET("/list", goods.CategoryList)
			goodscategoryGrp.GET("/info", goods.CategoryInfo)
		}

		shopGrp := domainGroup.Group("/shop")
		{
			shopGrp.GET("/info", shop.Info)
		}

		areaGrp := domainGroup.Group("/area")
		{
			areaGrp.GET("/list", cache.CachePage(store, 10*time.Minute, area.List))
			areaGrp.GET("/info", area.Info)
		}

		cartGrp := domainGroup.Group("/cart", wechat.WechatAccess())
		{
			cartGrp.POST("/list", cart.List)
			cartGrp.POST("/create", cart.Create)
			cartGrp.POST("/update", cart.Update)
			cartGrp.POST("/del", cart.Del)

			cartGrp.GET("/exist", cart.Exist)
			cartGrp.GET("/info", cart.Info)
			cartGrp.POST("/check", cart.Check)
			cartGrp.GET("/totalNum", cart.TotalNum)

		}

		addressGrp := domainGroup.Group("/address", wechat.WechatAccess())
		{
			addressGrp.POST("/setDefault", address.SetDefault)
			addressGrp.GET("/default", address.Default)
			addressGrp.GET("/list", address.List)
			addressGrp.GET("/info", address.Info)
			addressGrp.POST("/create", address.Create)
			addressGrp.POST("/update", address.Update)
			addressGrp.POST("/del", address.Del)
		}

		orderGrp := domainGroup.Group("/order", wechat.WechatAccess())
		{
			orderGrp.GET("/stateNum", order.StateNum)
			orderGrp.GET("/list", order.List)
			orderGrp.GET("/info", order.Info)
			orderGrp.POST("/cancel", order.Cancel)
			orderGrp.POST("/confirmReceipt", order.ConfirmReceipt)
			orderGrp.POST("/logistics", order.Logistics) // 物流查询
			orderGrp.GET("/goodsList", order.GoodsList)
			orderGrp.GET("/goodsInfo", order.GoodsInfo)

		}

		userGrp := domainGroup.Group("/user")
		{
			userGrp.POST("/login", user.WechatLogin)
			userGrp.POST("/register", user.WechatRegister)
			userGrp.GET("/self", wechat.WechatAccess(), user.WechatSelf)
		}

		buyGrp := domainGroup.Group("buy", wechat.WechatAccess())
		{
			buyGrp.POST("/calculate", buy.Calculate)
			buyGrp.POST("/create", buy.Create)
			buyGrp.POST("/pay", buy.Pay)
		}

		collectGrp := domainGroup.Group("collect", wechat.WechatAccess())
		{
			collectGrp.GET("/info", collect.Info)
			collectGrp.POST("/create", collect.Create)
			collectGrp.POST("/del", collect.Del)
		}

		callbackGrp := domainGroup.Group("/callback")
		{
			callbackGrp.POST("/wechatMiniNotify", buy.WechatMiniNotify)
		}
	}
	return r
}

func InitStat() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome ServerApi Stats",
			},
		)
	})

	return r
}
