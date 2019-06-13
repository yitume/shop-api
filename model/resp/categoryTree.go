package resp

type GoodsCategoryList struct {
	Id       int                 `json:"id"`
	Pid      int                 `json:"pid"`
	Name     string              `json:"name"`
	Children []GoodsCategoryList `json:"children"`
	Icon     string              `json:"icon"`
	Sort     int                 `json:"-"`
}
