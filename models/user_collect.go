package models

import "time"

// UserCollectModel 自定义第三张表  记录用户什么时候收藏了什么文章
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32"`
	CreatedAt time.Time
}

func (m *UserCollectModel) TableName() string {
	// 自定义表名的逻辑
	return "user_collect"
}
