package mysql

type WechatAutoReplyKeywords struct {
	AutoReplyId int    `json:"auto_reply_id" form:"auto_reply_id" ` //
	Key         string `json:"key" form:"key" `                     //
	MatchMode   string `json:"match_mode" form:"match_mode" `       //
	OpenId      int    `json:"open_id" form:"open_id" `             // 商家ID

}

func (*WechatAutoReplyKeywords) TableName() string {
	return "wechat_auto_reply_keywords"
}
