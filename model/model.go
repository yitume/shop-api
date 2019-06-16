package model

import (
	"github.com/jinzhu/gorm"
	"github.com/yitume/muses/pkg/cache/redis"
	"github.com/yitume/muses/pkg/database/mysql"
	"github.com/yitume/muses/pkg/logger"
)

var (
	Logger *logger.Client
	Db     *gorm.DB
	Redigo *redis.Client
)

func Init() {
	Db = mysql.Caller("mall")
	Logger = logger.Caller("system")
	Redigo = redis.Caller("auth")
}
