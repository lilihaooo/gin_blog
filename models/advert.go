package models

// AdvertModel 广告表
type AdvertModel struct {
	Model
	Title  string `json:"title" gorm:"size:32;comment:显示的标题"`       // 显示的标题
	Href   string `json:"href" gorm:"comment:跳转链接"`                 // 跳转链接
	Images string `json:"images" gorm:"comment:图片"`                 // 图片
	IsShow bool   `json:"is_show" gorm:"comment:是否展示 1: 展示 0: 不展示"` // 是否展示
}

func (m *AdvertModel) TableName() string {
	// 自定义表名的逻辑
	return "advert"
}
