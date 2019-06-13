package collect

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/yitume/shop-api/router/mdw/wechat"
	"github.com/yitume/shop-api/service"
)

func Info(c *gin.Context) {
	uid := wechat.WechatUid(c)
	_, respArr := service.GoodsCollect.ListPage(c, mysql.Conds{
		"uid": uid,
	}, trans.ReqPage{
		Current:  0,
		PageSize: 1000,
		Sort:     "id asc",
	})

	goodsIs := make([]int, 0)
	for _, value := range respArr {
		goodsIs = append(goodsIs, value.GoodsId)
	}

	cnt, output := service.Goods.ListPage(c, mysql.Conds{
		"id": mysql.Cond{"in", goodsIs},
	}, trans.ReqPage{
		Current:  0,
		PageSize: 1000,
		Sort:     "id asc",
	})
	base.JSONList(c, output, cnt)
}

func Create(c *gin.Context) {
	req := trans.ReqCollectCreate{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	var err error
	tx := model.Db.Begin()
	err = service.GoodsCollect.DeleteX(c, tx, mysql.Conds{
		"uid":      uid,
		"goods_id": req.GoodsId,
	})

	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	_, err = service.Goods.InfoX(c, mysql.Conds{
		"id": req.GoodsId,
	})
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	createInfo := &mysql.GoodsCollect{
		Uid:        uid,
		GoodsId:    req.GoodsId,
		CreateTime: time.Now().Unix(),
		OpenId:     openId,
	}
	err = service.GoodsCollect.Create(c, tx, createInfo)
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	base.JSON(c, base.MsgOk)
}

func Del(c *gin.Context) {
	req := trans.ReqCollectDel{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)
	var err error
	tx := model.Db.Begin()
	err = service.GoodsCollect.DeleteX(c, tx, mysql.Conds{
		"uid":      uid,
		"goods_id": req.GoodsId,
	})

	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}
	base.JSON(c, base.MsgOk)
}
