package model

import (
	"github.com/jinzhu/gorm"
	cgorm "github.com/yitume/caller/gorm"
	"github.com/yitume/caller/redigo"
	"github.com/yitume/caller/zap"
)

var (
	Logger *zap.ZapClient
	Db     *gorm.DB
	Redigo *redigo.RedigoClient
)

func Init() {
	Db = cgorm.Caller("mall").DB
	Logger = zap.Caller("system")
	Redigo = redigo.Caller("auth")
}
