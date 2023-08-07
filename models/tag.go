package models

// TagModel 标签表
type TagModel struct {
	Model
	Title string `gorm:"size:16" json:"title"` // 标签的名称
}

func (m *TagModel) TableName() string {
	// 自定义表名的逻辑
	return "tag"
}
