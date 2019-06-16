package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"go.uber.org/zap"
)

type shop struct{}

func InitShop() *shop {
	return &shop{}
}

// Create 新增一条记录
func (*shop) Create(c *gin.Context, db *gorm.DB, data *mysql.Shop) (err error) {

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create shop create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*shop) Update(c *gin.Context, db *gorm.DB, paramOpenId int, ups mysql.Ups) (err error) {
	var sql = "`open_id`=?"
	var binds = []interface{}{paramOpenId}

	if err = db.Table("shop").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("shop update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*shop) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)

	if err = db.Table("shop").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("shop update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*shop) Delete(c *gin.Context, db *gorm.DB, paramOpenId int) (err error) {
	var sql = "`open_id`=?"
	var binds = []interface{}{paramOpenId}

	if err = db.Table("shop").Where(sql, binds...).Update(mysql.Ups{"delete_time": time.Now().Unix()}).Error; err != nil {
		model.Logger.Error("shop delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*shop) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)

	if err = db.Table("shop").Where(sql, binds...).Update(mysql.Ups{"delete_time": time.Now().Unix()}).Error; err != nil {
		model.Logger.Error("shop delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*shop) Info(c *gin.Context, paramOpenId int) (resp mysql.Shop, err error) {
	var sql = "`open_id`=?"
	var binds = []interface{}{paramOpenId}

	sql += " AND delete_time = 0"
	if err = model.Db.Table("shop").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("shop info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*shop) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.Shop, err error) {
	sql, binds := mysql.BuildQuery(conds)

	sql += " AND delete_time = 0"
	if err = model.Db.Table("shop").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("shop info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*shop) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.Shop, err error) {
	sql, binds := mysql.BuildQuery(conds)

	sql += " AND delete_time = 0"
	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("shop").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("shop info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*shop) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.Shop, err error) {
	sql, binds := mysql.BuildQuery(conds)

	sql += " AND delete_time = 0"
	mysqlSlice := make([]mysql.Shop, 0)
	resp = make(map[int]mysql.Shop, 0)
	if err = model.Db.Table("shop").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("shop info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.OpenId] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*shop) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.Shop) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)

	sql += " AND delete_time = 0"
	db := model.Db.Table("shop").Where(sql, binds...)
	respList = make([]mysql.Shop, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
