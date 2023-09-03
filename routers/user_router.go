package routers

import (
	v1 "blog_gin/api/v1"
	"github.com/gin-gonic/gin"
)

func UserRouter(appGroup *gin.RouterGroup) {
	menuGroup := appGroup.Group("v1")
	menuGroup.GET("/user", v1.ApiGroupApp.UserApi.UserList)
	menuGroup.PUT("/user_role", v1.ApiGroupApp.UserApi.UserUpdateRole)
	menuGroup.PUT("/password", v1.ApiGroupApp.UserApi.UserUpdatePassword)
	menuGroup.DELETE("/logout", v1.ApiGroupApp.UserApi.Logout)
	menuGroup.POST("/user/send_email", v1.ApiGroupApp.UserApi.BindEmailSend)
	menuGroup.POST("/user/validate_email", v1.ApiGroupApp.UserApi.BindEmailValidate)
}
