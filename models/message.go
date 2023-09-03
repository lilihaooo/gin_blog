package models

import "blog_gin/global"

// MessageModel 记录消息
type MessageModel struct {
	Model
	UserID uint
	RoomID uint
	Data   string
}

func (m *MessageModel) TableName() string {
	// 自定义表名的逻辑
	return "message"
}

func (m *MessageModel) AddOneMessage(massage *MessageModel) error {
	return global.DB.Save(massage).Error
}
