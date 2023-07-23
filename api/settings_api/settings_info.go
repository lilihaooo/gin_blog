package settings_api

import (
	"blog_gin/pkg/app"
	"blog_gin/pkg/constant/error_const"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	user := User{
		"lihao",
		18,
	}

	appG := app.Gin{C: c}
	appG.Response(error_const.SUCCESS, "", user)
}
