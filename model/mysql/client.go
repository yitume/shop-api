package mysql

type Client struct {
	Aid         string
	Secret      string
	Extra       string
	RedirectUri string
}

// Set User's table name to be `profiles`
func (Client) TableName() string {
	return "app"
}
