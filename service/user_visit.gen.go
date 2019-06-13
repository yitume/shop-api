package service

import (
	"time"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type userVisit struct{}

func InitUserVisit() *userVisit {
	return &userVisit{}
}

// Create 新增一条记录
func (*userVisit) Create(c *gin.Context, db *gorm.DB, data *mysql.UserVisit) (err error) {
	data.OpenId = mdw.OpenId(c)
	data.CreateTime = time.Now().Unix()
	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create userVisit create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*userVisit) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_visit").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("user_visit update error", zap.String("err", err.Error()))
		return
	}
	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*userVisit) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_visit").Where(sql, binds...).Delete(&mysql.UserVisit{}).Error; err != nil {
		model.Logger.Error("user_visit delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*userVisit) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.UserVisit, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("user_visit").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("user_visit info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*userVisit) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.UserVisit, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("user_visit").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("user_visit info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListPage 根据分页条件查询list
func (*userVisit) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.UserVisit) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("user_visit").Where(sql, binds...)
	respList = make([]mysql.UserVisit, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
