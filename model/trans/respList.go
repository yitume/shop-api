package trans

type RespList struct {
	List        interface{} `json:"list"`
	TotalNumber int         `json:"total_number"`
}
