package tree

import (
	"sort"

	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/resp"
)

type Store struct {
	pidMap map[int]map[int]resp.GoodsCategoryList
	idMap  map[int]map[int]resp.GoodsCategoryList
}

func NewTreeStore(list []mysql.GoodsCategory) (obj *Store) {
	obj = &Store{
		pidMap: make(map[int]map[int]resp.GoodsCategoryList),
		idMap:  make(map[int]map[int]resp.GoodsCategoryList),
	}

	for _, v := range list {
		if _, ok := obj.pidMap[v.Pid]; !ok {
			obj.pidMap[v.Pid] = make(map[int]resp.GoodsCategoryList)
		}
		if _, ok := obj.idMap[v.Id]; !ok {
			obj.idMap[v.Id] = make(map[int]resp.GoodsCategoryList)
		}

		storeData := resp.GoodsCategoryList{
			Id:       v.Id,
			Pid:      v.Pid,
			Name:     v.Name,
			Children: make([]resp.GoodsCategoryList, 0),
			Icon:     v.Icon,
			Sort:     v.Sort,
		}

		obj.pidMap[v.Pid][v.Id] = storeData
		obj.idMap[v.Id][v.Pid] = storeData
	}
	return
}

func (t *Store) ListToTree(root int) (tree []resp.GoodsCategoryList) {
	tree = make([]resp.GoodsCategoryList, 0)

	for id, item := range t.pidMap[root] {
		if t.pidMap[id] != nil {
			item.Children = t.ListToTree(id)
		}
		tree = append(tree, item)
	}
	sort.Slice(tree, func(i, j int) bool { return tree[i].Sort < tree[j].Sort })
	return
}
