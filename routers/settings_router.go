package routers

import (
	"blog_gin/api"
	"github.com/gin-gonic/gin"
)

func SettingsRouter(appGroup *gin.RouterGroup) {
	// 继续分组
	settingsGroup := appGroup.Group("settings")
	// settingsApi 接口
	settingsApi := api.ApiGroupApp.SettingsApi
	// 分组使用接口
	settingsGroup.GET("site_info", settingsApi.GetSiteInfo)
	settingsGroup.PUT("site_info", settingsApi.UpdateSiteInfo)

	settingsGroup.GET("email", settingsApi.GetEmail)
	settingsGroup.PUT("email", settingsApi.UpdateEmail)

	settingsGroup.GET("jwt", settingsApi.GetJwt)
	settingsGroup.PUT("jwt", settingsApi.UpdateJwt)

	settingsGroup.GET("qiniu", settingsApi.GetQiniu)
	settingsGroup.PUT("qiniu", settingsApi.UpdateQiniu)

	settingsGroup.GET("qq", settingsApi.GetQQ)
	settingsGroup.PUT("qq", settingsApi.UpdateQQ)
}
