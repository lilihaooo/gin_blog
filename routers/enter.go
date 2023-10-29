package routers

import (
	v1 "blog_gin/api/v1"
	"blog_gin/middleware/jwt"
	"blog_gin/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()
	r.POST("/login", v1.ApiGroupApp.UserApi.Login)
	// 需要登陆
	r.Use(jwt.JwtAuth())
	// 创建ws全局的连接
	//r.GET("ws_conn", service.AppService.WebsocketService.WebsocketConn)
	// 聊天服务
	r.GET("ws_chat", service.AppService.WebsocketService.WebsocketChat)

	apiGroup := r.Group("api")
	// 配置静态文件根目录  // todo 该代码能否去掉?
	uploadDir := "/uploads"
	r.Static("/image", uploadDir)

	// 系统设置api
	SettingsRouter(apiGroup)
	// 图片管理
	ImagesRouter(apiGroup)
	// 广告管理
	AdvertRouter(apiGroup)
	// 菜单管理
	MenuRouter(apiGroup)
	// 用户管理
	UserRouter(apiGroup)
	// 文章管理
	ArticleRouter(apiGroup)
	return r
}
