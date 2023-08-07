package routers

import (
	v1 "blog_gin/api/v1"
	"github.com/gin-gonic/gin"
)

func ImagesRouter(appGroup *gin.RouterGroup) {
	// v1ImagesApi 接口
	v1ImagesGroup := appGroup.Group("v1/image")
	v1ImagesApi := v1.ApiGroupApp.ImagesApi
	// 分组使用接口
	v1ImagesGroup.POST("", v1ImagesApi.ImageUpload)
	v1ImagesGroup.GET("", v1ImagesApi.ImagesList)
	v1ImagesGroup.DELETE("", v1ImagesApi.ImagesDelete)
	v1ImagesGroup.PUT("", v1ImagesApi.ImageUpdate)
}
