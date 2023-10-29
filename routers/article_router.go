package routers

import (
	v1 "blog_gin/api/v1"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(appGroup *gin.RouterGroup) {
	menuGroup := appGroup.Group("v1")
	menuGroup.POST("/article", v1.ApiGroupApp.ArticleApi.ArticleCreate)
	menuGroup.GET("/articles", v1.ApiGroupApp.ArticleApi.ArticleList)
	menuGroup.GET("/article", v1.ApiGroupApp.ArticleApi.ArticleDetail)
	menuGroup.GET("/article_by_title", v1.ApiGroupApp.ArticleApi.ArticleDetailByTitle)
	menuGroup.GET("/article_calendar", v1.ApiGroupApp.ArticleApi.ArticleCalendar)
	menuGroup.GET("/article_tag", v1.ApiGroupApp.ArticleApi.ArticleTag)
	menuGroup.PUT("/update_article_title", v1.ApiGroupApp.ArticleApi.ArticleUpdate)
}
