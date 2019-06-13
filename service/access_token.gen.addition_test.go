package service_test

import (
	"testing"

	"github.com/godefault/caller"
	"github.com/godefault/caller/ginsession"
	"github.com/godefault/caller/gorm"
	"github.com/godefault/caller/redigo"
	callerzap "github.com/godefault/caller/zap"

	"fmt"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/service"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	confFilePath := "../../conf/conf.toml"

	caller.Init(
		confFilePath,
		gorm.New,
		redigo.New,
		ginsession.New,
		callerzap.New,
	)

	model.Init()
	service.Init()
	service.InitGen()
}

func TestCreateAccessToken(t *testing.T) {
	Convey("创建access token", t, func() {
		resp, err := service.AccessToken.CreateAccessToken(1, 1500000000)
		fmt.Println(resp.AccessToken)
		So(resp.AccessToken, ShouldNotEqual, "")
		So(err, ShouldBeNil)
	})

}

func TestCheckAccessToken(t *testing.T) {
	Convey("创建access token", t, func() {
		resp, err := service.AccessToken.CreateAccessToken(1, 1500000000)
		fmt.Println(resp.AccessToken)
		So(resp.AccessToken, ShouldNotEqual, "")
		So(err, ShouldBeNil)

		flag := service.FaAccessToken.CheckAccessToken(resp.AccessToken)
		fmt.Println(flag)
	})

}
