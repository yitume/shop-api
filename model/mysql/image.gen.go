package mysql

type Image struct {
	Id         int     `json:"id" form:"id" gorm:"primary_key"` //
	Name       string  `json:"name" form:"name" `               //
	Size       float64 `json:"size" form:"size" `               //
	Type       string  `json:"type" form:"type" `               //
	Url        string  `json:"url" form:"url" `                 //
	CreateTime int64   `json:"create_time" form:"create_time" ` //
	OpenId     int     `json:"open_id" form:"open_id" `         // 商家ID

}

func (*Image) TableName() string {
	return "image"
}
