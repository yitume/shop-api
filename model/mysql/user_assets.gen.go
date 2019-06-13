package mysql

type UserAssets struct {
	Id         int     `json:"id" form:"id" gorm:"primary_key"` // ID
	Uid        int     `json:"uid" form:"uid" gorm:"index"`     // 用户ID
	Points     int     `json:"points" form:"points" `           // 积分
	Balance    float64 `json:"balance" form:"balance" `         // 余额
	DeleteTime int64   `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int     `json:"open_id" form:"open_id" `         // 商家ID

}

func (*UserAssets) TableName() string {
	return "user_assets"
}
