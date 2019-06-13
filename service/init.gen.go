package service

var (
	Upload                  *upload
	Plugin                  *plugin
	MessageState            *messageState
	Page                    *page
	DiscountGoods           *discountGoods
	PdCash                  *pdCash
	OrderPay                *orderPay
	OrderRefundReason       *orderRefundReason
	Goods                   *goods
	Sms                     *sms
	Extend                  *extend
	UserOpen                *userOpen
	OffpayArea              *offpayArea
	Area                    *area
	WechatUser              *wechatUser
	Group                   *group
	OrderExtend             *orderExtend
	Express                 *express
	Queue                   *queue
	GoodsSpecValue          *goodsSpecValue
	Article                 *article
	GoodsSku                *goodsSku
	Address                 *address
	Coupon                  *coupon
	Message                 *message
	OrderStatis             *orderStatis
	GoodsCategory           *goodsCategory
	UserTemp                *userTemp
	Shipper                 *shipper
	AuthRule                *authRule
	Order                   *order
	Image                   *image
	TransportExtend         *transportExtend
	Discount                *discount
	WechatAutoReply         *wechatAutoReply
	SaleNum                 *saleNum
	UserPointsLog           *userPointsLog
	PdLog                   *pdLog
	OrderLog                *orderLog
	UserLevel               *userLevel
	Freight                 *freight
	AccessToken             *accessToken
	UserAlias               *userAlias
	GoodsSpec               *goodsSpec
	Setting                 *setting
	Transport               *transport
	Visit                   *visit
	OrderRefund             *orderRefund
	Version                 *version
	GroupConfig             *groupConfig
	GoodsEvaluate           *goodsEvaluate
	GoodsImage              *goodsImage
	AuthGroup               *authGroup
	OrderGoods              *orderGoods
	Biz                     *biz
	GoodsCart               *goodsCart
	MessageType             *messageType
	PdRecharge              *pdRecharge
	Shop                    *shop
	System                  *system
	GoodsCollect            *goodsCollect
	SendArea                *sendArea
	FullcutGoods            *fullcutGoods
	PayLog                  *payLog
	WechatAutoReplyKeywords *wechatAutoReplyKeywords
	GroupGoods              *groupGoods
	Fullcut                 *fullcut
	VerifyCode              *verifyCode
	SmsProvider             *smsProvider
	CouponGoods             *couponGoods
	Cart                    *cart
	Invoice                 *invoice
	SmsScene                *smsScene
	InfoCategory            *infoCategory
	Wechat                  *wechat
	Material                *material
	UserVisit               *userVisit
	Dispatch                *dispatch
	UserProfile             *userProfile
	Cron                    *cron
	GoodsCategoryIds        *goodsCategoryIds
	UserAssets              *userAssets
	UserAccount             *userAccount
	OrderRefundLog          *orderRefundLog
	AuthGroupAccess         *authGroupAccess
	WechatBroadcast         *wechatBroadcast
	Fd                      *fd
)

func InitGen() {
	Sms = InitSms()
	Extend = InitExtend()
	Goods = InitGoods()
	Group = InitGroup()
	OrderExtend = InitOrderExtend()
	Express = InitExpress()
	Queue = InitQueue()
	UserOpen = InitUserOpen()
	OffpayArea = InitOffpayArea()
	Area = InitArea()
	WechatUser = InitWechatUser()
	Article = InitArticle()
	GoodsSku = InitGoodsSku()
	Address = InitAddress()
	GoodsSpecValue = InitGoodsSpecValue()
	UserTemp = InitUserTemp()
	Shipper = InitShipper()
	AuthRule = InitAuthRule()
	Order = InitOrder()
	Coupon = InitCoupon()
	Message = InitMessage()
	OrderStatis = InitOrderStatis()
	GoodsCategory = InitGoodsCategory()
	TransportExtend = InitTransportExtend()
	Discount = InitDiscount()
	Image = InitImage()
	PdLog = InitPdLog()
	OrderLog = InitOrderLog()
	UserLevel = InitUserLevel()
	Freight = InitFreight()
	WechatAutoReply = InitWechatAutoReply()
	SaleNum = InitSaleNum()
	UserPointsLog = InitUserPointsLog()
	UserAlias = InitUserAlias()
	GoodsSpec = InitGoodsSpec()
	Setting = InitSetting()
	Transport = InitTransport()
	AccessToken = InitAccessToken()
	OrderRefund = InitOrderRefund()
	Version = InitVersion()
	GroupConfig = InitGroupConfig()
	Visit = InitVisit()
	AuthGroup = InitAuthGroup()
	OrderGoods = InitOrderGoods()
	Biz = InitBiz()
	GoodsCart = InitGoodsCart()
	GoodsEvaluate = InitGoodsEvaluate()
	GoodsImage = InitGoodsImage()
	PdRecharge = InitPdRecharge()
	Shop = InitShop()
	System = InitSystem()
	GoodsCollect = InitGoodsCollect()
	MessageType = InitMessageType()
	PayLog = InitPayLog()
	WechatAutoReplyKeywords = InitWechatAutoReplyKeywords()
	GroupGoods = InitGroupGoods()
	Fullcut = InitFullcut()
	SendArea = InitSendArea()
	FullcutGoods = InitFullcutGoods()
	SmsProvider = InitSmsProvider()
	CouponGoods = InitCouponGoods()
	VerifyCode = InitVerifyCode()
	Invoice = InitInvoice()
	SmsScene = InitSmsScene()
	Cart = InitCart()
	Wechat = InitWechat()
	Material = InitMaterial()
	InfoCategory = InitInfoCategory()
	UserProfile = InitUserProfile()
	Cron = InitCron()
	GoodsCategoryIds = InitGoodsCategoryIds()
	UserAssets = InitUserAssets()
	UserVisit = InitUserVisit()
	Dispatch = InitDispatch()
	OrderRefundLog = InitOrderRefundLog()
	AuthGroupAccess = InitAuthGroupAccess()
	WechatBroadcast = InitWechatBroadcast()
	Fd = InitFd()
	UserAccount = InitUserAccount()
	DiscountGoods = InitDiscountGoods()
	PdCash = InitPdCash()
	OrderPay = InitOrderPay()
	OrderRefundReason = InitOrderRefundReason()
	Upload = InitUpload()
	Plugin = InitPlugin()
	MessageState = InitMessageState()
	Page = InitPage()

}
