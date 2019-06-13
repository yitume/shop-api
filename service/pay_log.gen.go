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

type payLog struct{}

func InitPayLog() *payLog {
	return &payLog{}
}

// Create 新增一条记录
func (*payLog) Create(c *gin.Context, db *gorm.DB, data *mysql.PayLog) (err error) {
	data.OpenId = mdw.OpenId(c)

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create payLog create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*payLog) Update(c *gin.Context, db *gorm.DB, paramLogId int, ups mysql.Ups) (err error) {
	var sql = "`log_id`=?"
	var binds = []interface{}{paramLogId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("pay_log").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("pay_log update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*payLog) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("pay_log").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("pay_log update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*payLog) Delete(c *gin.Context, db *gorm.DB, paramLogId int) (err error) {
	var sql = "`log_id`=?"
	var binds = []interface{}{paramLogId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("pay_log").Where(sql, binds...).Update(mysql.Ups{"delete_time": time.Now().Unix()}).Error; err != nil {
		model.Logger.Error("pay_log delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*payLog) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("pay_log").Where(sql, binds...).Update(mysql.Ups{"delete_time": time.Now().Unix()}).Error; err != nil {
		model.Logger.Error("pay_log delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*payLog) Info(c *gin.Context, paramLogId int) (resp mysql.PayLog, err error) {
	var sql = "`log_id`=?"
	var binds = []interface{}{paramLogId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	if err = model.Db.Table("pay_log").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("pay_log info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*payLog) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.PayLog, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	if err = model.Db.Table("pay_log").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("pay_log info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*payLog) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.PayLog, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("pay_log").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("pay_log info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*payLog) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.PayLog, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	mysqlSlice := make([]mysql.PayLog, 0)
	resp = make(map[int]mysql.PayLog, 0)
	if err = model.Db.Table("pay_log").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("pay_log info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.LogId] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*payLog) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.PayLog) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	db := model.Db.Table("pay_log").Where(sql, binds...)
	respList = make([]mysql.PayLog, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
