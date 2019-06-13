package trans

type ReqAreaListInfo struct {
	Pid   int `form:"pid"`
	Level int `form:"level"`
}

type ReqAreaInfo struct {
	Name string `form:"name"`
}
