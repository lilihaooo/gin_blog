package settings_api

import (
	"blog_gin/config"
	"blog_gin/core"
	"blog_gin/global"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) GetEmail(c *gin.Context) {
	email := global.Config.Email
	res.OkWithData(c, email)
}

func (SettingsApi) UpdateEmail(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 只将这一部分修改
	global.Config.Email = cr
	if !core.SetYaml() {
		res.Fail(c, res.FAIL_OPER, "")
		return
	}
	res.Ok(c)
}
