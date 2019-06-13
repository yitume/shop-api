package mysql

type Shop struct {
	Name                         string `json:"name" form:"name" `                                                         // 店铺名字
	Logo                         string `json:"logo" form:"logo" `                                                         // 店铺标志图片
	ContactNumber                string `json:"contact_number" form:"contact_number" `                                     // 联系电话
	Description                  string `json:"description" form:"description" `                                           // 店铺描述
	ColorScheme                  int    `json:"color_scheme" form:"color_scheme" `                                         // 配色方案
	PortalTemplateId             int    `json:"portal_template_id" form:"portal_template_id" `                             // 店铺首页模板id
	WechatPlatformQr             string `json:"wechat_platform_qr" form:"wechat_platform_qr" `                             // 微信公众平台二维码
	GoodsCategoryStyle           int    `json:"goods_category_style" form:"goods_category_style" `                         // 店铺分类页风格
	DeleteTime                   int64  `json:"delete_time" form:"delete_time" `                                           // 软删除时间
	Salt                         string `json:"salt" form:"salt" `                                                         // 盐
	Host                         string `json:"host" form:"host" `                                                         //
	OrderAutoCloseExpires        int    `json:"order_auto_close_expires" form:"order_auto_close_expires" `                 // 待付款订单N秒后自动关闭订单
	OrderAutoConfirmExpires      int    `json:"order_auto_confirm_expires" form:"order_auto_confirm_expires" `             // 已发货订单后自动确认收货
	OrderAutoCloseRefoundExpires int    `json:"order_auto_close_refound_expires" form:"order_auto_close_refound_expires" ` // 已收货订单后关闭退款／退货功能，0代表确认收货后无法维权
	OpenId                       int    `json:"open_id" form:"open_id" gorm:"primary_key"`                                 //

}

func (*Shop) TableName() string {
	return "shop"
}
