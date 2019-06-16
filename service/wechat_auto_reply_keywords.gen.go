package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/router/mdw"
	"go.uber.org/zap"
)

type wechatAutoReplyKeywords struct{}

func InitWechatAutoReplyKeywords() *wechatAutoReplyKeywords {
	return &wechatAutoReplyKeywords{}
}

// Create 新增一条记录
func (*wechatAutoReplyKeywords) Create(c *gin.Context, db *gorm.DB, data *mysql.WechatAutoReplyKeywords) (err error) {
	data.OpenId = mdw.OpenId(c)

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create wechatAutoReplyKeywords create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*wechatAutoReplyKeywords) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("wechat_auto_reply_keywords").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("wechat_auto_reply_keywords update error", zap.String("err", err.Error()))
		return
	}
	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*wechatAutoReplyKeywords) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("wechat_auto_reply_keywords").Where(sql, binds...).Delete(&mysql.WechatAutoReplyKeywords{}).Error; err != nil {
		model.Logger.Error("wechat_auto_reply_keywords delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*wechatAutoReplyKeywords) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.WechatAutoReplyKeywords, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("wechat_auto_reply_keywords").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("wechat_auto_reply_keywords info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*wechatAutoReplyKeywords) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.WechatAutoReplyKeywords, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("wechat_auto_reply_keywords").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("wechat_auto_reply_keywords info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListPage 根据分页条件查询list
func (*wechatAutoReplyKeywords) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.WechatAutoReplyKeywords) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("wechat_auto_reply_keywords").Where(sql, binds...)
	respList = make([]mysql.WechatAutoReplyKeywords, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
