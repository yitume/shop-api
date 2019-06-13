package mysql

type Authorize struct {
	Client      string
	Code        string
	ExpiresIn   int32
	Scope       string
	RedirectUri string
	State       string
	Extra       string
	CreatedAt   int64
}

func (Authorize) TableName() string {
	return "authorize"
}
