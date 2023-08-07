package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()
	apiGroup := r.Group("api")

	// 配置静态文件根目录
	uploadDir := "/uploads"
	r.Static("/image", uploadDir)

	// 系统设置api
	SettingsRouter(apiGroup)
	ImagesRouter(apiGroup)

	// 图片管理

	return r
}
