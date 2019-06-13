package mysql

type Expires struct {
	Id        uint32
	Token     string
	ExpiresAt int64
}

func (Expires) TableName() string {
	return "expires"
}
