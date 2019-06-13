package address

import (
	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/resp"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/pkg/common/code"
	"github.com/yitume/shop-api/router/base"
	"github.com/yitume/shop-api/router/mdw"
	"github.com/yitume/shop-api/router/mdw/wechat"
	"github.com/yitume/shop-api/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	uid := wechat.WechatUid(c)
	cnt, list := service.Address.ListPage(c, mysql.Conds{
		"uid": uid,
	}, trans.ReqPage{
		Current:  1,
		PageSize: 20,
		Sort:     "id desc",
	})

	output := make([]resp.AddressInfo, 0)
	for _, info := range list {
		output = append(output, resp.AddressInfo{
			Id:            info.Id,
			Truename:      info.Truename,
			ProvinceId:    info.ProvinceId,
			CityId:        info.CityId,
			AreaId:        info.AreaId,
			Address:       info.Address,
			CombineDetail: info.CombineDetail,
			MobilePhone:   info.MobilePhone,
			IsDefault:     info.IsDefault,
			Type:          info.Type,
		})
	}
	base.JSONList(c, output, cnt)
}

func Info(c *gin.Context) {
	req := trans.ReqAddressInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgAddressInfoReqError)
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	info, err := service.Address.InfoX(c, mysql.Conds{
		"id":      req.Id,
		"uid":     uid,
		"open_id": openId,
	})

	if err != nil {
		base.JSON(c, code.MsgAddressInfoGetError, "error")
		return
	}
	base.JSON(c, code.MsgOk, resp.AddressInfo{
		Id:            info.Id,
		Truename:      info.Truename,
		ProvinceId:    info.ProvinceId,
		CityId:        info.CityId,
		AreaId:        info.AreaId,
		Address:       info.Address,
		CombineDetail: info.CombineDetail,
		MobilePhone:   info.MobilePhone,
		IsDefault:     info.IsDefault,
		Type:          info.Type,
	})
}

func Create(c *gin.Context) {
	req := trans.ReqAddressCreateInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, code.MsgAddressCreateReqError)
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)

	area, err := service.Area.Info(c, req.AreaId)
	if err != nil {
		base.JSON(c, code.MsgAddressCreateGetInfoError1)
		return
	}
	city, err := service.Area.Info(c, area.Pid)
	if err != nil {
		base.JSON(c, code.MsgAddressCreateGetInfoError2)
		return
	}
	province, err := service.Area.Info(c, city.Pid)
	if err != nil {
		base.JSON(c, code.MsgAddressCreateGetInfoError3)
		return
	}

	tx := model.Db.Begin()
	if req.IsDefault == 1 {
		err = service.Address.UpdateX(c, tx, mysql.Conds{
			"uid":     uid,
			"open_id": openId,
		}, mysql.Ups{
			"is_default": 0,
		})
		if err != nil {
			tx.Rollback()
			base.JSON(c, code.MsgAddressCreateError1)
			return
		}
	}

	createAddress := &mysql.Address{
		Id:            0,
		Uid:           uid,
		Truename:      req.Truename,
		ProvinceId:    province.Id,
		CityId:        city.Id,
		AreaId:        area.Id,
		Address:       req.Address,
		CombineDetail: province.Name + "-" + city.Name + "-" + area.Name,
		TelPhone:      "",
		MobilePhone:   req.MobilePhone,
		ZipCode:       "",
		IsDefault:     req.IsDefault,
		Type:          req.Type,
		StreetId:      0,
		DeleteTime:    0,
		OpenId:        openId,
	}

	err = service.Address.Create(c, tx, createAddress)

	if err != nil {
		tx.Rollback()
		base.JSON(c, code.MsgAddressCreateError2)
		return
	}
	tx.Commit()

	base.JSON(c, base.MsgOk)
}

func SetDefault(c *gin.Context) {
	req := trans.ReqAddressSetDefault{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	err := service.Address.UpdateX(c, model.Db, mysql.Conds{
		"id":      req.Id,
		"uid":     uid,
		"open_id": openId,
	}, mysql.Ups{
		"is_default": 1,
	})
	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	base.JSON(c, base.MsgOk)
}

func Default(c *gin.Context) {
	uid := wechat.WechatUid(c)

	info, _ := service.Address.InfoX(c, mysql.Conds{
		"uid":        uid,
		"is_default": 1,
	})

	base.JSON(c, base.MsgOk, resp.AddressInfo{
		Id:            info.Id,
		Truename:      info.Truename,
		ProvinceId:    info.ProvinceId,
		CityId:        info.CityId,
		AreaId:        info.AreaId,
		Address:       info.Address,
		CombineDetail: info.CombineDetail,
		MobilePhone:   info.MobilePhone,
		IsDefault:     info.IsDefault,
		Type:          info.Type,
	})
}

func Update(c *gin.Context) {
	req := trans.ReqAddressUpdateInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)

	area, err := service.Area.Info(c, req.AreaId)
	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	city, err := service.Area.Info(c, area.Pid)
	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	province, err := service.Area.Info(c, city.Pid)
	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	tx := model.Db.Begin()
	if req.IsDefault == 1 {
		err = service.Address.UpdateX(c, tx, mysql.Conds{
			"uid":     uid,
			"open_id": openId,
		}, mysql.Ups{
			"is_default": 0,
		})
		if err != nil {
			tx.Rollback()
			base.JSON(c, base.MsgErr)
			return
		}
	}

	err = service.Address.UpdateX(c, tx, mysql.Conds{
		"id":      req.Id,
		"open_id": openId,
	}, mysql.Ups{
		"truename":       req.Truename,
		"province_id":    province.Id,
		"city_id":        city.Id,
		"area_id":        area.Id,
		"address":        req.Address,
		"combine_detail": province.Name + "-" + city.Name + "-" + area.Name,
		"tel_phone":      "",
		"mobile_phone":   req.MobilePhone,
		"is_default":     req.IsDefault,
		"type":           req.Type,
	})

	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}
	tx.Commit()

	base.JSON(c, base.MsgOk)
}

func Del(c *gin.Context) {
	req := trans.ReqAddressDel{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	err := service.Cart.DeleteX(c, model.Db, mysql.Conds{
		"uid":     uid,
		"open_id": openId,
		"id":      req.Id,
	})

	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	base.JSON(c, base.MsgOk)

}
