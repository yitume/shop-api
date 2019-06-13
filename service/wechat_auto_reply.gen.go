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

type wechatAutoReply struct{}

func InitWechatAutoReply() *wechatAutoReply {
	return &wechatAutoReply{}
}

// Create 新增一条记录
func (*wechatAutoReply) Create(c *gin.Context, db *gorm.DB, data *mysql.WechatAutoReply) (err error) {
	data.OpenId = mdw.OpenId(c)
	data.CreateTime = time.Now().Unix()
	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create wechatAutoReply create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*wechatAutoReply) Update(c *gin.Context, db *gorm.DB, paramId int, ups mysql.Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("wechat_auto_reply").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("wechat_auto_reply update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*wechatAutoReply) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("wechat_auto_reply").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("wechat_auto_reply update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*wechatAutoReply) Delete(c *gin.Context, db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("wechat_auto_reply").Where(sql, binds...).Delete(&mysql.WechatAutoReply{}).Error; err != nil {
		model.Logger.Error("wechat_auto_reply delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*wechatAutoReply) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("wechat_auto_reply").Where(sql, binds...).Delete(&mysql.WechatAutoReply{}).Error; err != nil {
		model.Logger.Error("wechat_auto_reply delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*wechatAutoReply) Info(c *gin.Context, paramId int) (resp mysql.WechatAutoReply, err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("wechat_auto_reply").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("wechat_auto_reply info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*wechatAutoReply) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.WechatAutoReply, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("wechat_auto_reply").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("wechat_auto_reply info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*wechatAutoReply) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.WechatAutoReply, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("wechat_auto_reply").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("wechat_auto_reply info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*wechatAutoReply) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.WechatAutoReply, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	mysqlSlice := make([]mysql.WechatAutoReply, 0)
	resp = make(map[int]mysql.WechatAutoReply, 0)
	if err = model.Db.Table("wechat_auto_reply").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("wechat_auto_reply info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*wechatAutoReply) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.WechatAutoReply) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("wechat_auto_reply").Where(sql, binds...)
	respList = make([]mysql.WechatAutoReply, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
