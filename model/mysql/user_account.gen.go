package mysql

import "database/sql"

type UserAccount struct {
	Uid                 int             `json:"uid" form:"uid" gorm:"primary_key"`                 //
	AvailablePredeposit sql.NullFloat64 `json:"available_predeposit" form:"available_predeposit" ` //
	FreezePredeposit    sql.NullFloat64 `json:"freeze_predeposit" form:"freeze_predeposit" `       //
	OpenId              int             `json:"open_id" form:"open_id" `                           // 商家ID

}

func (*UserAccount) TableName() string {
	return "user_account"
}
