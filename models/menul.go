package models

import (
	"blog_gin/models/ctype"
)

// MenuModel 菜单表  菜单的路径可以是 /path 也可以是路由别名
type MenuModel struct {
	Model
	Title        string        `gorm:"size:32;comment:标题" json:"title"`                                                                         // 标题
	Path         string        `gorm:"size:32;comment:路径" json:"path"`                                                                          // 路径
	Slogan       string        `gorm:"size:64;comment:slogan" json:"slogan"`                                                                      // slogan
	Abstract     ctype.Array   `gorm:"type:string;comment:简介" json:"abstract"`                                                                  // 简介
	AbstractTime int           `gorm:"comment:简介的切换时间" json:"abstract_time"`                                                               // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner;joinForeignKey:MenuID;JoinReferences:BannerID;comment:菜单的图片列表" json:"banners"` // 菜单的图片列表
	BannerTime   int           `gorm:"comment:菜单图片的切换时间" json:"banner_time"`                                                             // 菜单图片的切换时间 为 0 表示不切换
	Sort         int           `gorm:"size:10;comment:菜单的顺序" json:"sort"`                                                                    // 菜单的顺序
}

func (m *MenuModel) TableName() string {
	// 自定义表名的逻辑
	return "menu"
}
