package routers

import (
	v1 "blog_gin/api/v1"
	"github.com/gin-gonic/gin"
)

func ImagesRouter(appGroup *gin.RouterGroup) {
	// v1ImagesApi 接口
	v1ImagesGroup := appGroup.Group("v1/image")
	// 分组使用接口
	v1ImagesGroup.POST("", v1.ApiGroupApp.ImagesApi.ImageUpload)
	v1ImagesGroup.GET("", v1.ApiGroupApp.ImagesApi.ImagesList)
	v1ImagesGroup.DELETE("", v1.ApiGroupApp.ImagesApi.ImagesDelete)
	v1ImagesGroup.PUT("", v1.ApiGroupApp.ImagesApi.ImageUpdate)
}
