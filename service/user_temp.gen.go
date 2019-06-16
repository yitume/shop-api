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

type userTemp struct{}

func InitUserTemp() *userTemp {
	return &userTemp{}
}

// Create 新增一条记录
func (*userTemp) Create(c *gin.Context, db *gorm.DB, data *mysql.UserTemp) (err error) {
	data.OpenId = mdw.OpenId(c)

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create userTemp create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*userTemp) Update(c *gin.Context, db *gorm.DB, paramUserId int, ups mysql.Ups) (err error) {
	var sql = "`user_id`=?"
	var binds = []interface{}{paramUserId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_temp").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("user_temp update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*userTemp) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_temp").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("user_temp update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*userTemp) Delete(c *gin.Context, db *gorm.DB, paramUserId int) (err error) {
	var sql = "`user_id`=?"
	var binds = []interface{}{paramUserId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_temp").Where(sql, binds...).Delete(&mysql.UserTemp{}).Error; err != nil {
		model.Logger.Error("user_temp delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*userTemp) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_temp").Where(sql, binds...).Delete(&mysql.UserTemp{}).Error; err != nil {
		model.Logger.Error("user_temp delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*userTemp) Info(c *gin.Context, paramUserId int) (resp mysql.UserTemp, err error) {
	var sql = "`user_id`=?"
	var binds = []interface{}{paramUserId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("user_temp").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("user_temp info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*userTemp) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.UserTemp, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("user_temp").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("user_temp info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*userTemp) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.UserTemp, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("user_temp").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("user_temp info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*userTemp) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.UserTemp, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	mysqlSlice := make([]mysql.UserTemp, 0)
	resp = make(map[int]mysql.UserTemp, 0)
	if err = model.Db.Table("user_temp").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("user_temp info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.UserId] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*userTemp) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.UserTemp) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("user_temp").Where(sql, binds...)
	respList = make([]mysql.UserTemp, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
