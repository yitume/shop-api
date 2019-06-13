package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type GoodsEvaluate struct {
	Id                int                               `json:"id" form:"id" gorm:"primary_key"`               // 评价ID
	OrderId           int                               `json:"order_id" form:"order_id" `                     // 订单表自增ID
	OrderNo           int64                             `json:"order_no" form:"order_no" `                     // 订单编号
	OrderGoodsId      int                               `json:"order_goods_id" form:"order_goods_id" `         // 订单商品表编号
	GoodsId           int                               `json:"goods_id" form:"goods_id" `                     // 商品主表id
	GoodsSkuId        int                               `json:"goods_sku_id" form:"goods_sku_id" `             // 商品表编号
	GoodsTitle        string                            `json:"goods_title" form:"goods_title" `               // 商品名称
	GoodsPrice        float64                           `json:"goods_price" form:"goods_price" `               // 商品价格
	GoodsImg          string                            `json:"goods_img" form:"goods_img" `                   // 商品图片
	Score             int                               `json:"score" form:"score" `                           // 1-5分
	Content           string                            `json:"content" form:"content" `                       // 信誉评价内容，有可能会存表情
	IsAnonymous       int                               `json:"is_anonymous" form:"is_anonymous" `             // 0表示不是 1表示是匿名评价
	CreateTime        int64                             `json:"create_time" form:"create_time" `               // 评价时间
	UserId            int                               `json:"user_id" form:"user_id" `                       // 评价人编号
	State             int                               `json:"state" form:"state" `                           // 评价信息的状态 0为正常 1为禁止显示
	Remark            string                            `json:"remark" form:"remark" `                         // 管理员对评价的处理备注
	Explain           string                            `json:"explain" form:"explain" `                       // 解释内容
	Images            GoodsEvaluateImagesJson           `json:"images" form:"images" `                         // 晒单图片json
	AdditionalContent string                            `json:"additional_content" form:"additional_content" ` // 追加评论
	AdditionalTime    int64                             `json:"additional_time" form:"additional_time" `       // 追加时间
	AdditionalImages  GoodsEvaluateAdditionalImagesJson `json:"additional_images" form:"additional_images" `   // 追加评价图片
	ReplyContent      string                            `json:"reply_content" form:"reply_content" `           // 回复评价的内容
	ReplyTime         int64                             `json:"reply_time" form:"reply_time" `                 // 回复时间
	ReplyContent2     string                            `json:"reply_content2" form:"reply_content2" `         // 回复追加的内容
	ReplyTime2        int64                             `json:"reply_time2" form:"reply_time2" `               // 回复时间
	Display           int                               `json:"display" form:"display" `                       // 回复状态 1是显示 0是隐藏
	Top               int                               `json:"top" form:"top" `                               // 置顶 0：不是；1：是；
	DeleteTime        int64                             `json:"delete_time" form:"delete_time" `               // 软删除时间
	OpenId            int                               `json:"open_id" form:"open_id" `                       // 商户id

}

func (*GoodsEvaluate) TableName() string {
	return "goods_evaluate"
}

// 请在model/mysql/addtion.go里定义GoodsEvaluateImagesJson结构体
func (c GoodsEvaluateImagesJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsEvaluateImagesJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 请在model/mysql/addtion.go里定义GoodsEvaluateAdditionalImagesJson结构体
func (c GoodsEvaluateAdditionalImagesJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsEvaluateAdditionalImagesJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
