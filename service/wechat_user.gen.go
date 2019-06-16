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

type wechatUser struct{}

func InitWechatUser() *wechatUser {
	return &wechatUser{}
}

// Create 新增一条记录
func (*wechatUser) Create(c *gin.Context, db *gorm.DB, data *mysql.WechatUser) (err error) {

	data.CreateTime = time.Now().Unix()
	data.UpdateTime = time.Now().Unix()
	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create wechatUser create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*wechatUser) Update(c *gin.Context, db *gorm.DB, paramId int, ups mysql.Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	ups["update_time"] = time.Now().Unix()
	if err = db.Table("wechat_user").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("wechat_user update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*wechatUser) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {
	ups["update_time"] = time.Now().Unix()
	sql, binds := mysql.BuildQuery(conds)

	if err = db.Table("wechat_user").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("wechat_user update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*wechatUser) Delete(c *gin.Context, db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("wechat_user").Where(sql, binds...).Delete(&mysql.WechatUser{}).Error; err != nil {
		model.Logger.Error("wechat_user delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*wechatUser) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)

	if err = db.Table("wechat_user").Where(sql, binds...).Delete(&mysql.WechatUser{}).Error; err != nil {
		model.Logger.Error("wechat_user delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*wechatUser) Info(c *gin.Context, paramId int) (resp mysql.WechatUser, err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = model.Db.Table("wechat_user").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("wechat_user info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*wechatUser) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.WechatUser, err error) {
	sql, binds := mysql.BuildQuery(conds)

	if err = model.Db.Table("wechat_user").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("wechat_user info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*wechatUser) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.WechatUser, err error) {
	sql, binds := mysql.BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("wechat_user").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("wechat_user info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*wechatUser) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.WechatUser, err error) {
	sql, binds := mysql.BuildQuery(conds)

	mysqlSlice := make([]mysql.WechatUser, 0)
	resp = make(map[int]mysql.WechatUser, 0)
	if err = model.Db.Table("wechat_user").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("wechat_user info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*wechatUser) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.WechatUser) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)

	db := model.Db.Table("wechat_user").Where(sql, binds...)
	respList = make([]mysql.WechatUser, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
