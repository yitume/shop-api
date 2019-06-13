package trans

type Banner struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	ReferId    uint32 `json:"referId"`
	LinkUrl    string `json:"linkUrl"`
	PicUrl     string `json:"picUrl"`
	Remark     string `json:"remark"`
	Title      string `json:"title"`
	BannerType int    `json:"bannerType"`
}

func (Banner) TableName() string {
	return "banner"
}

type NoticeList struct {
	ID   uint32 `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (NoticeList) TableName() string {
	return "notice"
}

type NoticeInfo struct {
	ID      uint32 `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (NoticeInfo) TableName() string {
	return "notice"
}
