package mysql

type UserVisit struct {
	UserId     int   `json:"user_id" form:"user_id" `         //
	CreateTime int64 `json:"create_time" form:"create_time" ` //
	OpenId     int   `json:"open_id" form:"open_id" `         // 商家ID

}

func (*UserVisit) TableName() string {
	return "user_visit"
}
