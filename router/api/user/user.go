package user

import (
	"time"

	"github.com/yitume/shop-api/pkg/bootstrap"
	"github.com/yitume/shop-api/pkg/common/code"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/medivhzhan/weapp"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/pkg/random"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/yitume/shop-api/router/mdw/wechat"
	"github.com/yitume/shop-api/service"
)

// todo 前端 有时序性问题。需要调整
func WechatLogin(c *gin.Context) {
	req := trans.ReqUserLogin{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgUserAccountLoginReqError)
		return
	}

	// 1 登录微信
	respWechat, err := weapp.Login(bootstrap.Conf.Wechat.AppID, bootstrap.Conf.Wechat.AppSecret, req.WechatMiniCode)
	if err != nil {
		base.JSON(c, code.MsgUserAccountLoginWechatError)
		return
	}

	m := map[string]string{
		"sessionKey": respWechat.SessionKey,
		"openId":     respWechat.OpenID,
	}

	// 2 设置60s有效期
	err = model.Redigo.Hmset(req.WechatMiniCode, m, 10)
	if err != nil {
		base.JSON(c, code.MsgAccountLoginWechatRedisError)
		return
	}

	// 3 获取数据库信息
	respMsql, err := service.UserOpen.InfoX(c, mysql.Conds{
		"wechat_openid": respWechat.OpenID,
	})

	if err != nil {
		base.JSON(c, code.MsgAccountLoginWechatMysqlError, req.WechatMiniCode)
		return
	}

	// 3 校验用户是否存在
	if respMsql.Uid == 0 {
		base.JSON(c, code.MsgAccountLoginWechatUserNotExist)
		return
	}

	// 4 创建token
	respToken, err := service.AccessToken.CreateAccessToken(c, respMsql.Uid, time.Now().Unix())
	if err != nil {
		base.JSON(c, code.MsgAccountLoginWechatCreateTokenError)
		return
	}

	base.JSON(c, code.MsgOk, respToken)
	return
}

func WechatRegister(c *gin.Context) {
	req := trans.ReqUserReigster{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgAccountRegisterReqError)
		return
	}
	// todo hmget
	sessionKey, err := model.Redigo.HgetString(req.WechatMiniParam.Code, "sessionKey")
	userInfo, err := weapp.DecryptUserInfo(req.WechatMiniParam.RawData, req.WechatMiniParam.EncryptedData, req.WechatMiniParam.Signature, req.WechatMiniParam.Iv, sessionKey)
	if err != nil {
		base.JSON(c, code.MsgAccountRegisterRedisError)
		return
	}

	// 3 获取数据库信息

	respMsql, err := service.UserOpen.InfoX(c, mysql.Conds{
		"wechat_openid": userInfo.OpenID,
	})

	if err != nil && err != gorm.ErrRecordNotFound || respMsql.Uid > 0 {
		base.JSON(c, code.MsgAccountRegisterMysqlGetError)
		return
	}
	// union id todo 需要加上union id判断，以后可能有用

	userName := "wechat_mini_" + userInfo.OpenID + random.GetRandomString(8)

	createUserInfo := mysql.UserOpenInfoAggregateJson{
		OpenID:   userInfo.OpenID,
		Nickname: userInfo.Nickname,
		Gender:   userInfo.Gender,
		Province: userInfo.Province,
		Language: userInfo.Language,
		Country:  userInfo.Country,
		City:     userInfo.City,
		Avatar:   userInfo.Avatar,
		UnionID:  userInfo.UnionID,
		Watermark: mysql.UserOpenInfoAggregateJsonWatermark{
			AppID:     userInfo.Watermark.AppID,
			Timestamp: userInfo.Watermark.Timestamp,
		},
	}

	createUserOpen := &mysql.UserOpen{
		Genre:         1,
		Name:          userName,
		WechatOpenid:  userInfo.OpenID,
		AppOpenid:     "",
		MiniOpenid:    userInfo.OpenID,
		Unionid:       userInfo.OpenID,
		AccessToken:   "",
		ExpiresIn:     0,
		RefreshToken:  "",
		Scope:         "",
		Nickname:      userInfo.Nickname,
		Avatar:        userInfo.Avatar,
		Sex:           userInfo.Gender,
		Country:       userInfo.Country,
		Province:      userInfo.Province,
		City:          userInfo.City,
		InfoAggregate: createUserInfo,
		State:         0,
		CreateTime:    time.Now().Unix(),
		DeleteTime:    0,
		OpenId:        mdw.OpenId(c),
	}
	err = service.UserOpen.Create(c, model.Db, createUserOpen)
	if err != nil {
		base.JSON(c, code.MsgAccountRegisterMysqlCreateUserError)
		return
	}

	// 4 创建token
	respToken, err := service.AccessToken.CreateAccessToken(c, createUserOpen.Uid, time.Now().Unix())
	if err != nil {
		base.JSON(c, code.MsgAccountRegisterMysqlCreateTokenError)
		return
	}

	base.JSON(c, base.MsgOk, respToken)
}

func WechatSelf(c *gin.Context) {
	uid := wechat.WechatUid(c)
	// todo openid
	userInfo, _ := service.UserOpen.Info(c, uid)

	assetsInfo, _ := service.UserAssets.InfoX(c, mysql.Conds{
		"uid": uid,
	})

	base.JSON(c, base.MsgOk, map[string]interface{}{
		"uid":      uid,
		"nickname": userInfo.Nickname,
		"name":     userInfo.Name,
		"avatar":   userInfo.Avatar,
		"assets":   assetsInfo,
	})

}
