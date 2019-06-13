package service

import (
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type userWechat struct{}

func InitUserWechat() *userWechat {
	return &userWechat{}
}

// Create 新增一条记录
func (*userWechat) Create(db *gorm.DB, data *mysql.UserWechat) (err error) {

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create userWechat create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*userWechat) Update(db *gorm.DB, paramUserId int, ups mysql.Ups) (err error) {

	if err = db.Table("user_wechat").Where("`user_id`=?", paramUserId).Updates(ups).Error; err != nil {
		model.Logger.Error("user_wechat update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*userWechat) UpdateX(db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	if err = db.Table("user_wechat").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("user_wechat update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*userWechat) Delete(db *gorm.DB, paramUserId int) (err error) {

	if err = db.Table("user_wechat").Where("`user_id`=?", paramUserId).Delete(&mysql.UserWechat{}).Error; err != nil {
		model.Logger.Error("user_wechat delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*userWechat) DeleteX(db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)

	if err = db.Table("user_wechat").Where(sql, binds...).Delete(&mysql.UserWechat{}).Error; err != nil {
		model.Logger.Error("user_wechat delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*userWechat) Info(paramUserId int) (resp mysql.UserWechat, err error) {
	var sql = "`user_id`=?"
	var binds = []interface{}{paramUserId}

	if err = model.Db.Table("user_wechat").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("user_wechat info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*userWechat) InfoX(conds mysql.Conds) (resp mysql.UserWechat, err error) {
	sql, binds := mysql.BuildQuery(conds)

	if err = model.Db.Table("user_wechat").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("user_wechat info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*userWechat) List(conds mysql.Conds, extra ...string) (resp []mysql.UserWechat, err error) {
	sql, binds := mysql.BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("user_wechat").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("user_wechat info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*userWechat) ListMap(conds mysql.Conds) (resp map[int]mysql.UserWechat, err error) {
	sql, binds := mysql.BuildQuery(conds)

	mysqlSlice := make([]mysql.UserWechat, 0)
	resp = make(map[int]mysql.UserWechat, 0)
	if err = model.Db.Table("user_wechat").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("user_wechat info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.UserId] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*userWechat) ListPage(conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.UserWechat) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)

	db := model.Db.Table("user_wechat").Where(sql, binds...)
	respList = make([]mysql.UserWechat, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
