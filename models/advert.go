package models

import "gorm.io/gorm"

// AdvertModel 广告表
type AdvertModel struct {
	gorm.Model
	Title  string `gorm:"size:32;comment:显示的标题" json:"title"` // 显示的标题
	Href   string `json:"href;comment:跳转链接"`                  // 跳转链接
	Images string `json:"images;comment:图片"`                  // 图片
	IsShow bool   `json:"is_show;comment:是否展示"`               // 是否展示
}

func (m *AdvertModel) TableName() string {
	// 自定义表名的逻辑
	return "advert"
}
