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

type groupConfig struct{}

func InitGroupConfig() *groupConfig {
	return &groupConfig{}
}

// Create 新增一条记录
func (*groupConfig) Create(c *gin.Context, db *gorm.DB, data *mysql.GroupConfig) (err error) {
	data.OpenId = mdw.OpenId(c)

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create groupConfig create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*groupConfig) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("group_config").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("group_config update error", zap.String("err", err.Error()))
		return
	}
	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*groupConfig) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("group_config").Where(sql, binds...).Delete(&mysql.GroupConfig{}).Error; err != nil {
		model.Logger.Error("group_config delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*groupConfig) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.GroupConfig, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("group_config").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("group_config info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*groupConfig) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.GroupConfig, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("group_config").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("group_config info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListPage 根据分页条件查询list
func (*groupConfig) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.GroupConfig) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("group_config").Where(sql, binds...)
	respList = make([]mysql.GroupConfig, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
