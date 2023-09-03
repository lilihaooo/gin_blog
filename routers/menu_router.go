package routers

import (
	v1 "blog_gin/api/v1"
	"github.com/gin-gonic/gin"
)

func MenuRouter(appGroup *gin.RouterGroup) {
	menuGroup := appGroup.Group("v1/menu")
	menuGroup.POST("", v1.ApiGroupApp.MenuApi.MenuCreate)
	menuGroup.GET("", v1.ApiGroupApp.MenuApi.MenuList)
	menuGroup.GET("/name", v1.ApiGroupApp.MenuApi.MenuNameList)
	menuGroup.PUT("", v1.ApiGroupApp.MenuApi.MenuUpdate)
	menuGroup.DELETE("", v1.ApiGroupApp.MenuApi.MenuDelete)
	menuGroup.GET("/:id", v1.ApiGroupApp.MenuApi.MenuDetail)
}
