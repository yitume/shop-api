package mysql

type Refresh struct {
	Token  string
	Access string
}

func (Refresh) TableName() string {
	return "refresh"
}
