package mysql

type Access struct {
	Client       string
	Authorize    string
	Previous     string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int32
	Scope        string
	RedirectUri  string
	Extra        string
	CreatedAt    int64
}

func (Access) TableName() string {
	return "access"
}
