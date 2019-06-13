package mysql

import (
	"database/sql/driver"
	"encoding/json"
)

type OrduerExtendReciverInfoJson struct {
}

type GoodsBodyJson []struct {
	Type string `json:"type"`
	Val  struct {
		Content string `json:"content"`
	} `json:"value"`
}

type GoodsSpecList struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	ValueList []CreateSpecValue `json:"value_list"`
}

type GoodsSpecListJson []struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	ValueList []CreateSpecValue `json:"value_list"`
}

type CreateSpecValue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GoodsSkuSpecJson []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ValueID   int    `json:"value_id"`
	ValueName string `json:"value_name"`
	ValueImg  string `json:"value_img"`
}

type IntsJson []int

func (c IntsJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *IntsJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type StringsJson []string

func (c StringsJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *StringsJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type GoodsSkusJson []GoodsSku

func (c GoodsSkusJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *GoodsSkusJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type FreightAreasJson []FreightAreas

type FreightAreas struct {
	AreaIds          []int
	FirstAmount      float64
	FirstFee         float64
	AdditionalAmount float64
	AdditionalFee    float64
}

type FullcutHierarchyJson []string

type GoodsImageSpecImagesJson StringsJson

type GoodsEvaluateImagesJson StringsJson

type GoodsEvaluateAdditionalImagesJson StringsJson

type MaterialMediaJson StringsJson

type OrderExtendOrderGoodsJson []Goods

type OrderExtendReciverInfoJson struct {
	Name          string `json:"name"`
	CombineDetail string `json:"combine_detail"`
	Phone         string `json:"phone"`
	Type          string `json:"type"`
	Address       string `json:"address"`
}

type OrderExtendInvoiceInfoJson StringsJson

type OrderGoodsGoodsSpecJson []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ValueID   int    `json:"value_id"`
	ValueName string `json:"value_name"`
	ValueImg  string `json:"value_img"`
}

type OrderRefundUserImagesJson StringsJson

type OrderRefundTrackingImagesJson StringsJson

type OrderRefundGoodsSpecJson []Goods

type PageBodyJson string

type SendAreaAreaIdsJson IntsJson

type SettingConfigJson struct {
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}

type SmsProviderConfigJson string

type UserOpenInfoAggregateJson struct {
	OpenID    string                             `json:"openId"`
	Nickname  string                             `json:"nickName"`
	Gender    int                                `json:"gender"`
	Province  string                             `json:"province"`
	Language  string                             `json:"language"`
	Country   string                             `json:"country"`
	City      string                             `json:"city"`
	Avatar    string                             `json:"avatarUrl"`
	UnionID   string                             `json:"unionId"`
	Watermark UserOpenInfoAggregateJsonWatermark `json:"watermark"`
}

type UserOpenInfoAggregateJsonWatermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type UserWechatTagidListJson struct {
}

type WechatAutoReplySubscribeReplayContentJson struct{}

type WechatAutoReplyReplyContentJson struct{}

type WechatAutoReplyKeysJson struct{}

type WechatBroadcastConditionJson struct{}
type WechatBroadcastSendContentJson struct{}
type WechatBroadcastOpenidsJson struct{}
type WechatUserTagidListJson struct{}
type AuthGroupRuleIdsJson []int
type GoodsImagesJson []string
type GoodsCategoryIdsJson []int
type GoodsSkuListJson []GoodsSku
