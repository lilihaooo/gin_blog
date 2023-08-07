package routers

import (
	"blog_gin/api/v1"
	v2 "blog_gin/api/v2"
	"github.com/gin-gonic/gin"
)

func SettingsRouter(appGroup *gin.RouterGroup) {
	// v1SettingsApi 接口
	v1SettingsGroup := appGroup.Group("v1/settings")
	v1SettingsApi := v1.ApiGroupApp.SettingsApi
	// 分组使用接口
	v1SettingsGroup.GET("site_info", v1SettingsApi.GetSiteInfo)
	v1SettingsGroup.PUT("site_info", v1SettingsApi.UpdateSiteInfo)

	v1SettingsGroup.GET("email", v1SettingsApi.GetEmail)
	v1SettingsGroup.PUT("email", v1SettingsApi.UpdateEmail)
	v1SettingsGroup.GET("jwt", v1SettingsApi.GetJwt)
	v1SettingsGroup.PUT("jwt", v1SettingsApi.UpdateJwt)

	v1SettingsGroup.GET("qiniu", v1SettingsApi.GetQiNiu)
	v1SettingsGroup.PUT("qiniu", v1SettingsApi.UpdateQiNiu)

	v1SettingsGroup.GET("qq", v1SettingsApi.GetQQ)
	v1SettingsGroup.PUT("qq", v1SettingsApi.UpdateQQ)

	// v2版本接口
	v2settingsGroup := appGroup.Group("v2/settings")
	// v1SettingsApi 接口
	v2SettingsApi := v2.ApiGroupApp.SettingsApi
	v2settingsGroup.GET("/:name", v2SettingsApi.GetSittingsInfo)
	v2settingsGroup.PUT("/:name", v2SettingsApi.UpdateSittingsInfo)

}
