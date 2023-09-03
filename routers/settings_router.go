package routers

import (
	"blog_gin/api/v1"
	v2 "blog_gin/api/v2"
	"github.com/gin-gonic/gin"
)

func SettingsRouter(appGroup *gin.RouterGroup) {
	// v1SettingsApi 接口
	v1SettingsGroup := appGroup.Group("v1/settings")

	// 分组使用接口
	v1SettingsGroup.GET("site_info", v1.ApiGroupApp.SettingsApi.GetSiteInfo)
	v1SettingsGroup.PUT("site_info", v1.ApiGroupApp.SettingsApi.UpdateSiteInfo)

	v1SettingsGroup.GET("email", v1.ApiGroupApp.SettingsApi.GetEmail)
	v1SettingsGroup.PUT("email", v1.ApiGroupApp.SettingsApi.UpdateEmail)
	v1SettingsGroup.GET("jwt", v1.ApiGroupApp.SettingsApi.GetJwt)
	v1SettingsGroup.PUT("jwt", v1.ApiGroupApp.SettingsApi.UpdateJwt)

	v1SettingsGroup.GET("qiniu", v1.ApiGroupApp.SettingsApi.GetQiNiu)
	v1SettingsGroup.PUT("qiniu", v1.ApiGroupApp.SettingsApi.UpdateQiNiu)

	v1SettingsGroup.GET("qq", v1.ApiGroupApp.SettingsApi.GetQQ)
	v1SettingsGroup.PUT("qq", v1.ApiGroupApp.SettingsApi.UpdateQQ)

	// v2版本接口
	v2settingsGroup := appGroup.Group("v2/settings")
	v2settingsGroup.GET("/:name", v2.ApiGroupApp.SettingsApi.GetSittingsInfo)
	v2settingsGroup.PUT("/:name", v2.ApiGroupApp.SettingsApi.UpdateSittingsInfo)

}
