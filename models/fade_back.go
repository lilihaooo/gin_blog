package models

import "gorm.io/gorm"

type FadeBackModel struct {
	gorm.Model
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ApplyContent string `gorm:"size:128;comment:回复的内容" json:"apply_content"` // 回复的内容
	IsApply      bool   `gorm:"comment:是否回复" json:"is_apply"`                // 是否回复
}

func (m *FadeBackModel) TableName() string {
	// 自定义表名的逻辑
	return "fade_back"
}
