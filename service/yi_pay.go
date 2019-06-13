package service

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yitume/shop-api/model"

	"github.com/yitume/shop-api/model/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type yiPay struct{}

// 变更预存款
func (*yiPay) ChangePrePay(c *gin.Context, changeType string, uid int, username string, amount float64, sn string, state int, adminName string) (err error) {

	logData := &mysql.PdLog{
		PdSn:       sn,
		Uid:        uid,
		Username:   username,
		Type:       changeType,
		CreateTime: time.Now().Unix(),
		State:      state,
		DeleteTime: 0,
	}
	updateUserAccount := make(mysql.Ups, 0)
	updateUserAccount["type"] = changeType

	switch changeType {
	case "order_type":
		logData.AvailableAmount = -amount
		logData.Remark = "下单，支付预存款，订单号: " + sn
		updateUserAccount["available_predposit"] = gorm.Expr("available_predposit-?", amount)
	case "order_freeze":
		logData.AvailableAmount = -amount
		logData.FreezeAmount = amount
		logData.Remark = "下单，冻结预存款，订单号: " + sn

		updateUserAccount["freeze_predeposit"] = gorm.Expr("freeze_predeposit+?", amount)
		updateUserAccount["available_predeposit"] = gorm.Expr("available_predeposit-?", amount)
	case "order_cancel":
		logData.AvailableAmount = amount
		logData.FreezeAmount = -amount
		logData.Remark = "取消订单，解冻预存款，订单号:" + sn

		updateUserAccount["freeze_predeposit"] = gorm.Expr("freeze_predeposit-?", amount)
		updateUserAccount["available_predeposit"] = gorm.Expr("available_predeposit+?", amount)
	case "order_comb_pay":
		logData.FreezeAmount = -amount
		logData.Remark = "下单，支付被冻结的预存款，订单号: " + sn

		updateUserAccount["freeze_predeposit"] = gorm.Expr("freeze_predeposit-?", amount)
	case "recharge":
		logData.AvailableAmount = amount
		logData.Remark = "充值，充值单号: " + sn
		logData.AdminName = adminName

		updateUserAccount["available_predeposit"] = gorm.Expr("available_predeposit+?", amount)
	case "refund":
		logData.AvailableAmount = amount
		logData.Remark = "确认退款，订单号: " + sn

		updateUserAccount["available_predeposit"] = gorm.Expr("available_predeposit+?", amount)

	case "cash_apply":
		logData.AvailableAmount = -amount
		logData.FreezeAmount = amount
		logData.Remark = "申请提现，冻结预存款，提现单号: " + sn

		updateUserAccount["available_predeposit"] = gorm.Expr("available_predeposit-?", amount)
		updateUserAccount["freeze_predeposit"] = gorm.Expr("freeze_predeposit+?", amount)
	case "cash_pay":
		logData.AvailableAmount = -amount
		logData.Remark = "提现成功，提现单号: " + sn
		logData.AdminName = adminName

		updateUserAccount["freeze_predeposit"] = gorm.Expr("freeze_predeposit-?", amount)
	case "cash_del":
		logData.AvailableAmount = amount
		logData.FreezeAmount = -amount
		logData.Remark = "取消提现申请，解冻预存款，提现单号: " + sn
		logData.AdminName = adminName

		updateUserAccount["available_predeposit"] = gorm.Expr("available_predeposit+?", amount)
		updateUserAccount["freeze_predeposit"] = gorm.Expr("freeze_predeposit-?", amount)
	default:
		err = errors.New("type is error")
		return
	}

	if changeType == "recharge" && logData.State == 0 {

	} else {
		// todo tans
		err = UserAccount.UpdateX(c, model.Db, mysql.Conds{
			"uid": uid,
		}, updateUserAccount)
		if err != nil {
			return
		}
	}

	err = PdLog.Create(c, model.Db, logData)
	return
}
