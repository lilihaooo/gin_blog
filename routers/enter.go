package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()
	apiGroup := r.Group("api")

	// 系统设置api
	SettingsRouter(apiGroup)
	return r
}
