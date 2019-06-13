package mysql

type Area struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` // ID 名称code
	Name       string `json:"name" form:"name" `               // 栏目名
	Pid        int    `json:"pid" form:"pid" `                 // 父栏目
	Location   string `json:"location" form:"location" `       // 全路径
	LevelName  string `json:"level_name" form:"level_name" `   // 级别名称
	Longitude  string `json:"longitude" form:"longitude" `     // 经度
	Latitude   string `json:"latitude" form:"latitude" `       // 维度
	Level      int    `json:"level" form:"level" `             // 级别
	Position   string `json:"position" form:"position" `       // 方位
	CityCode   string `json:"city_code" form:"city_code" `     // 城市代码
	ZipCode    string `json:"zip_code" form:"zip_code" `       // 邮编
	Pinyin     string `json:"pinyin" form:"pinyin" `           // 拼音
	Initial    string `json:"initial" form:"initial" `         // 首字母
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间

}

func (*Area) TableName() string {
	return "area"
}
