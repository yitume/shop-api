package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
)

const accessTokenIss = "github.com/yitume/shop-api"
const accessTokenExpireInterval = 7 * 24 * 60 * 60
const accessTokenKey = "yitumsK#xo"

type AccessTokenTicket struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (g *accessToken) CreateAccessToken(c *gin.Context, uid int, startTime int64) (resp AccessTokenTicket, err error) {
	accessTokenData := &mysql.AccessToken{
		Jti:        0,
		Sub:        uid,
		IaTime:     startTime,
		ExpTime:    startTime + accessTokenExpireInterval,
		Ip:         "",
		CreateTime: time.Now().Unix(),
		IsLogout:   0,
		IsInvalid:  0,
		LogoutTime: 0,
	}

	err = g.Create(c, model.Db, accessTokenData)
	if err != nil {
		return
	}
	tokenString, err := encodeAccessToken(accessTokenData.Jti, uid, startTime)
	if err != nil {
		return
	}
	resp.AccessToken = tokenString
	resp.ExpiresIn = accessTokenExpireInterval
	return
}

func (g *accessToken) CheckAccessToken(token string) bool {
	sc, err := g.DecodeAccessToken(token)
	if err != nil {
		model.Logger.Error("access_token CheckAccessToken error1", zap.String("err", err.Error()))
		return false
	}

	var resp mysql.AccessToken
	if err = model.Db.Table("access_token").Where("`jti`=? AND `sub`=? AND `exp_time`>=? AND `is_invalid`=? AND `is_logout`=?", sc["jti"], sc["sub"], sc["exp"], 0, 0).Find(&resp).Error; err != nil {
		model.Logger.Error("access_token CheckAccessToken error2", zap.String("err", err.Error()))
		return false
	}
	return true
}

func (g *accessToken) RefreshAccessToken(c *gin.Context, token string, startTime int64) (resp AccessTokenTicket, err error) {
	sc, err := g.DecodeAccessToken(token)
	if err != nil {
		model.Logger.Error("access_token CheckAccessToken error1", zap.String("err", err.Error()))
		return
	}

	refreshToken, err := encodeAccessToken(sc["jti"].(int), sc["uid"].(int), startTime)

	if err != nil {
		return
	}

	err = g.Update(c, model.Db, sc["jti"].(int), map[string]interface{}{
		"exp_time": startTime + accessTokenExpireInterval,
	})
	if err != nil {
		return
	}
	resp.AccessToken = refreshToken
	resp.ExpiresIn = accessTokenExpireInterval
	return
}

func encodeAccessToken(jwtId int, uid int, startTime int64) (tokenStr string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["jti"] = jwtId
	claims["iss"] = accessTokenIss
	claims["sub"] = uid
	claims["iat"] = startTime
	claims["exp"] = startTime + accessTokenExpireInterval
	token.Claims = claims

	tokenStr, err = token.SignedString([]byte(accessTokenKey))
	if err != nil {
		return
	}
	return
}

func (g *accessToken) DecodeAccessToken(token string) (resp jwt.MapClaims, err error) {
	tokenParse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessTokenKey), nil
	})
	if err != nil {
		return
	}
	var flag bool
	resp, flag = tokenParse.Claims.(jwt.MapClaims)
	if !flag {
		err = errors.New("assert error")
		return
	}
	return
}
