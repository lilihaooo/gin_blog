package models

import (
	"blog_gin/global"
	"blog_gin/models/ctype"
	"gorm.io/gorm"
	"os"
)

// BannerModel 横幅model
type BannerModel struct {
	gorm.Model
	Path      string          `gorm:"comment:图片路径" json:"path"`                          // 图片路径
	Hash      string          `gorm:"comment:图片的hash值，用于判断重复图片" json:"hash"`             // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38;comment:图片名称" json:"name"`                  // 图片名称
	ImageType ctype.ImageType `gorm:"default:1;comment:图片的类型， 本地还是七牛" json:"image_type"` // 图片的类型， 本地还是七牛
}

func (m *BannerModel) TableName() string {
	// 自定义表名的逻辑
	return "banner"
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Logrus.Error(err)
			return err
		}
	}
	return nil
}
