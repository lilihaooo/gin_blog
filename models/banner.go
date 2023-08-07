package models

import (
	"blog_gin/global"
	"blog_gin/models/ctype"
	"gorm.io/gorm"
	"os"
	"time"
)

// BannerModel 横幅model
type BannerModel struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	Path      string          `gorm:"comment:图片路径" json:"path"`                              // 图片路径
	Hash      string          `gorm:"comment:图片的hash值，用于判断重复图片" json:"hash"`         // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38;comment:图片名称" json:"name"`                      // 图片名称
	Type      ctype.ImageType `gorm:"comment:图片类型 1: 本地, 2: 服务器;default:1" json:"type"` // 图片类型 1: 本地, 2: 服务器
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (b *BannerModel) TableName() string {
	// 自定义表名的逻辑
	return "banner"
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.Type == ctype.Local {
		if err = os.Remove(b.Path); err != nil {
			global.Logrus.Error(err)
			return err
		}
	}
	return nil
}
