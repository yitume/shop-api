package code

const (
	MsgOk = 0

	// 10000 ~ 10199
	MsgUserAccountLoginReqError             = 10000
	MsgUserAccountLoginWechatError          = 10001
	MsgAccountLoginWechatRedisError         = 10002
	MsgAccountLoginWechatMysqlError         = 10003
	MsgAccountLoginWechatUserNotExist       = 10004
	MsgAccountLoginWechatCreateTokenError   = 10005
	MsgAccountRegisterReqError              = 10006
	MsgAccountRegisterRedisError            = 10007
	MsgAccountRegisterMysqlGetError         = 10008
	MsgAccountRegisterMysqlCreateUserError  = 10009
	MsgAccountRegisterMysqlCreateTokenError = 10010

	MsgGoodsInfoReqErr   = 11000
	MsgGoodsInfoMysqlErr = 11001
	MsgGoodsListReqErr   = 11002

	MsgCarInfoReqErr          = 12000
	MsgCarInfoEmptyErr        = 12001
	MsgCartCreateReqErr       = 12002
	MsgCartCreateCartExistErr = 12003

	// 20000 ~ 20200
	MsgAddressInfoReqError = 20000
	MsgAddressInfoGetError = 20001

	MsgAddressCreateReqError      = 20002
	MsgAddressCreateGetInfoError1 = 20003
	MsgAddressCreateGetInfoError2 = 20004
	MsgAddressCreateGetInfoError3 = 20005
	MsgAddressCreateError1        = 20006
	MsgAddressCreateError2        = 20007
)

var CodeMap = map[int]string{
	MsgOk: "ok",
	12003: "购物车已经存在",
}
