package mysql

type App struct {
	Aid         int    `json:"aid"`
	Name        string `json:"name"`
	Secret      string `json:"secret"`
	RedirectUri string `json:"redirectUri"`
	Extra       string `json:"extra"`
	CallNo      int    `json:"callNo"`
	Status      int    `json:"status"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

func (App) TableName() string {
	return "app"
}
