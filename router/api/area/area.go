package area

import (
	"github.com/gin-gonic/gin"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/model/mysql"
	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/pkg/area"
	"github.com/yitume/shop-api/router/base"
)

/**
 * 地区列表
 * @method GET
 * @param int pid 父级id
 * @param int level 层级 规则：设置显示下级行政区级数（行政区级别包括：省/直辖市、市、区/县3个级别）可选值：0、1、2  0：返回省/直辖市；1：返回省/直辖市、市；2：返回省/直辖市、市、区/县
 * @param int $tree 默认 0
 */
func List(c *gin.Context) {
	req := trans.ReqAreaListInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}
	areaList := make([]mysql.Area, 0)

	// todo 换成 service代码
	switch req.Level {
	case 0:
		model.Db.Table("area").Where(" level = ?  AND delete_time = ?", 1, 0).Find(&areaList)
	case 1:
		model.Db.Table("area").Where(" level in(1,2)  AND delete_time = ?", 0).Find(&areaList)
	case 2:
		// todo 东西太大了
		model.Db.Table("area").Where("level in(1,2,3)  AND delete_time = ?", 0).Find(&areaList)
	default:
		model.Db.Table("area").Where(" level = ?  AND delete_time = ?", 1, 0).Find(&areaList)

	}

	treeObj := area.InitTree(areaList)

	base.JSON(c, base.MsgOk, treeObj.ToTree(0))
}

/**
 * 区域信息
 * 区县级，包含省和市
 * @method GET
 * @param string $name     如：海珠区
 * @param string $up_level ,默认2，向上两层 @todo
 * @todo   拓展其他方式的获取 如id
 *                         为：不只为微信端做服务
 */

//todo
//
//combine_detail:`${result.items[0].name} ${result.items[1].name} ${result.items[2].name}`,
//area_id:result.items[2].id,
//truename:res.userName,
//mobile_phone:res.telNumber,
//address:res.detailInfo,

func Info(c *gin.Context) {
	req := trans.ReqAreaInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}
	area := mysql.Area{}
	model.Db.Select("id,pid,name").Where("name=?", req.Name).Find(&area)
	if area.Id == 0 {
		base.JSON(c, base.MsgErr, "request app list params is error")
		return
	}

	city := mysql.Area{}
	province := mysql.Area{}

	model.Db.Select("id,pid,name").Where("id=?", area.Pid).Find(&city)
	model.Db.Select("id,pid,name").Where("id=?", city.Pid).Find(&province)

	type test struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
		Pid  int    `json:"pid"`
	}

	output := []test{
		test{
			Name: province.Name,
			Id:   province.Id,
			Pid:  province.Pid,
		},
		test{
			Name: city.Name,
			Id:   city.Id,
			Pid:  city.Pid,
		},
		test{
			Name: area.Name,
			Id:   area.Id,
			Pid:  area.Pid,
		},
	}
	base.JSON(c, base.MsgOk, output)

}
