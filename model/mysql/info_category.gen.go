package mysql

type InfoCategory struct {
	Id             int    `json:"id" form:"id" gorm:"primary_key"`         //
	Title          string `json:"title" form:"title" `                     // 名字
	Pid            int    `json:"pid" form:"pid" `                         // 父级id
	Level          int    `json:"level" form:"level" `                     // 层级默认为0
	Sort           int    `json:"sort" form:"sort" `                       // 排序默认为0
	Status         int    `json:"status" form:"status" `                   // 状态 1为开启0为关闭
	Desc           string `json:"desc" form:"desc" `                       // 描述
	Name           string `json:"name" form:"name" `                       // 唯一标识
	TemplateIndex  string `json:"template_index" form:"template_index" `   // 频道页模板
	TemplateList   string `json:"template_list" form:"template_list" `     // 列表页模板
	TemplateDetail string `json:"template_detail" form:"template_detail" ` // 详情页模板
	TemplateType   int    `json:"template_type" form:"template_type" `     // 0列表1封面
	Img            string `json:"img" form:"img" `                         // 封面图
	Text           string `json:"text" form:"text" `                       // 详情
	Flags          string `json:"flags" form:"flags" `                     // 属性
	Url            string `json:"url" form:"url" `                         // 跳转链接
	Images         string `json:"images" form:"images" `                   // 多图上传
	Extend         string `json:"extend" form:"extend" `                   // 拓展
	Keywords       string `json:"keywords" form:"keywords" `               // 关键词
	ImgStatus      int    `json:"img_status" form:"img_status" `           // 1显示0不显示
	Img2           string `json:"img2" form:"img2" `                       //
	DeleteTime     int64  `json:"delete_time" form:"delete_time" `         // 软删除时间
	OpenId         int    `json:"open_id" form:"open_id" `                 // 商家ID

}

func (*InfoCategory) TableName() string {
	return "info_category"
}
