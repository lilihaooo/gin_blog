package v1

import (
	"blog_gin/api/v1/advert_api"
	"blog_gin/api/v1/article_api"
	"blog_gin/api/v1/images_api"
	"blog_gin/api/v1/menu_api"
	"blog_gin/api/v1/settings_api"
	"blog_gin/api/v1/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertsApi  advert_api.AdvertsApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	ArticleApi  article_api.ArticleApi
}

var ApiGroupApp = new(ApiGroup)
