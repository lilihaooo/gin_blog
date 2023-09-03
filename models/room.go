package models

type RoomModel struct {
	Model
	Number uint   // 房间号
	Name   string // 房间名称
	info   string // 简介
	UserID uint   `gorm:"foreignKey:UserID;comment:房主"` // 创建房间userID
}

func (m *RoomModel) TableName() string {
	// 自定义表名的逻辑
	return "room"
}
