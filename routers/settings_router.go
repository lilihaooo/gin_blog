package routers

import (
	"blog_gin/api"
)

func (r *RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi

	r.GET("settings", settingsApi.SettingsInfoView)
}
