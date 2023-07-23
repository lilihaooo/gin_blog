package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()

	apiRouterGroup := r.Group("api")

	routerGroup := RouterGroup{apiRouterGroup}
	// 系统设置api
	routerGroup.SettingsRouter()
	return r
}
