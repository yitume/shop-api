package mysql

type Fd struct {
	Id         int   `json:"id" form:"id" gorm:"primary_key"` //
	Fd         int   `json:"fd" form:"fd" `                   //
	Uid        int   `json:"uid" form:"uid" `                 //
	CreateTime int64 `json:"create_time" form:"create_time" ` //
	OpenId     int   `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Fd) TableName() string {
	return "fd"
}
