package models

import (
	"blog_gin/global"
)

// UserRoom 用户和房间的中间表
type UserRoomModel struct {
	Model
	RoomID uint `gorm:"foreignKey:RoomID"`
	UserID uint `gorm:"foreignKey:UserID"`
}

func (m *UserRoomModel) TableName() string {
	// 自定义表名的逻辑
	return "user_room"
}

func (m *UserRoomModel) UserInRoom(userID, RoomID uint) bool {
	if err := global.DB.Take(m, "user_id = ? and room_id = ?", userID, RoomID).Error; err != nil {
		return false
	}
	return true
}

func (m *UserRoomModel) GetRecordsWithCondition(condition string) (res []UserRoomModel, err error) {
	if err = global.DB.Where(condition).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil

}
