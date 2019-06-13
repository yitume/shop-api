package mysql

//
//CREATE TABLE `banner` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`refer_id` int(10) unsigned NOT NULL,
//`link_url` varchar(255) NOT NULL,
//`pic_url` varchar(255) NOT NULL,
//`remark` varchar(255) NOT NULL,
//`status` int(255) NOT NULL,
//`title` varchar(255) NOT NULL,
//`banner_type` int(255) NOT NULL,
//`created_by` int(255) NOT NULL,
//`updated_by` int(11) NOT NULL,
//`created_at` int(255) NOT NULL,
//`updated_at` int(11) NOT NULL,
//`sorted_num` int(11) NOT NULL,
//`open_id` int(10) unsigned NOT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;

type Banner struct {
	ID      uint32 `gorm:"primary_key" json:"id"`
	ReferId uint32
	OpenId  uint32
	LinkUrl string
	PicUrl  string
	Remark  string
	Title   string

	Status     int
	BannerType int

	CreatedBy int `json:"createdBy"`
	UpdatedBy int `json:"updatedBy"`

	CreatedAt int `json:"createdAt"`
	UpdatedAt int `json:"updatedAt"`
}

func (Banner) TableName() string {
	return "banner"
}
