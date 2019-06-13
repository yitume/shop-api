package mysql

type GoodsCategory struct {
	Id          int    `json:"id" form:"id" gorm:"primary_key"` //
	Name        string `json:"name" form:"name" `               // 分类名称
	Pid         int    `json:"pid" form:"pid" `                 // 父级id，如果不填为一级
	Icon        string `json:"icon" form:"icon" `               // 商品分类图标，图片地址
	Sort        int    `json:"sort" form:"sort" `               // 排序
	CreateTime  int64  `json:"create_time" form:"create_time" ` // 添加时间
	UpdateTime  int64  `json:"update_time" form:"update_time" ` // 更新时间
	Keywords    string `json:"keywords" form:"keywords" `       //
	Description string `json:"description" form:"description" ` //
	Grade       string `json:"grade" form:"grade" `             // 价格分级
	Img         string `json:"img" form:"img" `                 // 分类前面的小图标
	TypeId      int    `json:"type_id" form:"type_id" `         // 关联类型id
	Banner      string `json:"banner" form:"banner" `           // 横幅
	DeleteTime  int64  `json:"delete_time" form:"delete_time" ` // 软删除时间
	OpenId      int    `json:"open_id" form:"open_id" `         // 商户id

}

func (*GoodsCategory) TableName() string {
	return "goods_category"
}
