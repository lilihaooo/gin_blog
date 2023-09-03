package flag

import (
	"blog_gin/global"
	"blog_gin/models"
)

func Makemigrations() {
	var err error
	// SetupJoinTable方法只是为多对多关联的中间表进行设置
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
			&models.RoomModel{},
			&models.UserRoomModel{},
		)
	if err != nil {
		global.Logrus.Errorf("生成数据库表结构失败:%s", err)
		return
	}
	global.Logrus.Info("生成数据库表结构成功")
}
