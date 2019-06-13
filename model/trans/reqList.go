package trans

type ReqList struct {
	Page int    `form:"page"`
	Rows int    `form:"rows"`
	Sort string `form:"sort"`
}

type ReqInfo struct {
	Id int `form:"id"`
}
