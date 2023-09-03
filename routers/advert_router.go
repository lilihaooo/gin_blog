package routers

import (
	v1 "blog_gin/api/v1"
	"github.com/gin-gonic/gin"
)

func AdvertRouter(appGroup *gin.RouterGroup) {
	advertGroup := appGroup.Group("v1/advert")
	advertGroup.POST("", v1.ApiGroupApp.AdvertsApi.AdvertCreate)
	advertGroup.DELETE("", v1.ApiGroupApp.AdvertsApi.AdvertDelete)
	advertGroup.GET("", v1.ApiGroupApp.AdvertsApi.AdvertList)
	advertGroup.PUT("", v1.ApiGroupApp.AdvertsApi.AdvertUpdate)
}
