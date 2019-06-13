package mysql

type Article struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"` //
	CategoryId int    `json:"category_id" form:"category_id" ` // 分类id
	Title      string `json:"title" form:"title" `             // 名称
	Hits       int    `json:"hits" form:"hits" `               // 点击率
	Text       string `json:"text" form:"text" `               // 内容
	Desc       string `json:"desc" form:"desc" `               // 描述
	Sort       int    `json:"sort" form:"sort" `               // 排序
	Img        string `json:"img" form:"img" `                 // 图片
	CreateTime int64  `json:"create_time" form:"create_time" ` // 时间
	Url        string `json:"url" form:"url" `                 // 链接
	Status     int    `json:"status" form:"status" `           // 0未审核1审核
	Flags      string `json:"flags" form:"flags" `             // 自定义属性
	Source     string `json:"source" form:"source" `           // 来源
	Author     string `json:"author" form:"author" `           // 作者
	Keywords   string `json:"keywords" form:"keywords" `       // 关键字
	Images     string `json:"images" form:"images" `           // 多图
	Extend     string `json:"extend" form:"extend" `           // 多功能拓展 json
	Template   string `json:"template" form:"template" `       // 模板地址
	ImgStatus  int    `json:"img_status" form:"img_status" `   // 1显示0不显示
	Attachment string `json:"attachment" form:"attachment" `   // 附件
	DeleteTime int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId     int    `json:"open_id" form:"open_id" `         // 商户id

}

func (*Article) TableName() string {
	return "article"
}
