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

type authGroupAccess struct{}

func InitAuthGroupAccess() *authGroupAccess {
	return &authGroupAccess{}
}

// Create 新增一条记录
func (*authGroupAccess) Create(c *gin.Context, db *gorm.DB, data *mysql.AuthGroupAccess) (err error) {
	data.OpenId = mdw.OpenId(c)

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create authGroupAccess create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*authGroupAccess) Update(c *gin.Context, db *gorm.DB, paramGroupId int, ups mysql.Ups) (err error) {
	var sql = "`group_id`=?"
	var binds = []interface{}{paramGroupId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("auth_group_access").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("auth_group_access update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*authGroupAccess) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("auth_group_access").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("auth_group_access update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*authGroupAccess) Delete(c *gin.Context, db *gorm.DB, paramGroupId int) (err error) {
	var sql = "`group_id`=?"
	var binds = []interface{}{paramGroupId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("auth_group_access").Where(sql, binds...).Delete(&mysql.AuthGroupAccess{}).Error; err != nil {
		model.Logger.Error("auth_group_access delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*authGroupAccess) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("auth_group_access").Where(sql, binds...).Delete(&mysql.AuthGroupAccess{}).Error; err != nil {
		model.Logger.Error("auth_group_access delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*authGroupAccess) Info(c *gin.Context, paramGroupId int) (resp mysql.AuthGroupAccess, err error) {
	var sql = "`group_id`=?"
	var binds = []interface{}{paramGroupId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("auth_group_access").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("auth_group_access info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*authGroupAccess) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.AuthGroupAccess, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("auth_group_access").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("auth_group_access info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*authGroupAccess) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.AuthGroupAccess, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("auth_group_access").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("auth_group_access info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*authGroupAccess) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.AuthGroupAccess, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	mysqlSlice := make([]mysql.AuthGroupAccess, 0)
	resp = make(map[int]mysql.AuthGroupAccess, 0)
	if err = model.Db.Table("auth_group_access").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("auth_group_access info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.GroupId] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*authGroupAccess) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.AuthGroupAccess) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("auth_group_access").Where(sql, binds...)
	respList = make([]mysql.AuthGroupAccess, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
