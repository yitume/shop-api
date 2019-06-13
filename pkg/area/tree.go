package area

import (
	"sort"

	"github.com/yitume/shop-api/model/mysql"
)

type Area struct {
	Id        int    `json:"id" gorm:"primary_key"` // ID 名称code
	Name      string `json:"name" `                 // 栏目名
	Pid       int    `json:"pid" `                  // 父栏目
	Location  string `json:"location" `             // 全路径
	LevelName string `json:"level_name" `           // 级别名称
	Longitude string `json:"longitude" `            // 经度
	Latitude  string `json:"latitude" `             // 维度
	Level     int    `json:"level" `                // 级别
	Position  string `json:"position" `             // 方位
	CityCode  string `json:"city_code" `            // 城市代码
	ZipCode   string `json:"zip_code" `             // 邮编
	Pinyin    string `json:"pinyin" `               // 拼音
	Initial   string `json:"initial" `              // 首字母
	Children  []Area `json:"children"`
}

type Tree struct {
	itemMap map[int]Area
	// 					pid        id
	pidIdMap map[int]map[int]Area
	idPidMap map[int]map[int]Area
}

func InitTree(list []mysql.Area) *Tree {
	obj := &Tree{
		itemMap:  make(map[int]Area, 0),
		pidIdMap: make(map[int]map[int]Area),
		idPidMap: make(map[int]map[int]Area),
	}
	for _, v := range list {
		obj.itemMap[v.Id] = Area{
			Id:        v.Id,
			Name:      v.Name,
			Pid:       v.Pid,
			Location:  v.Location,
			LevelName: v.LevelName,
			Longitude: v.Longitude,
			Latitude:  v.Latitude,
			Level:     v.Level,
			Position:  v.Position,
			CityCode:  v.CityCode,
			ZipCode:   v.ZipCode,
			Pinyin:    v.Pinyin,
			Initial:   v.Initial,
			Children:  make([]Area, 0),
		}
	}
	for _, v := range obj.itemMap {
		if _, ok := obj.pidIdMap[v.Pid]; !ok {
			obj.pidIdMap[v.Pid] = make(map[int]Area)
		}
		if _, ok := obj.idPidMap[v.Id]; !ok {
			obj.idPidMap[v.Id] = make(map[int]Area)
		}
		obj.pidIdMap[v.Pid][v.Id] = v
		obj.idPidMap[v.Id][v.Pid] = v
	}
	return obj
}

func (t *Tree) ToTree(index int) []Area {
	tmp := make([]Area, 0)
	for id, item := range t.pidIdMap[index] {
		if t.pidIdMap[id] != nil {
			item.Children = t.ToTree(id)
		}
		tmp = append(tmp, item)
	}
	sort.Slice(tmp, func(i, j int) bool { return tmp[i].Level < tmp[j].Level })
	return tmp
}
