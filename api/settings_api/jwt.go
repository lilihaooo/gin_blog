package settings_api

import (
	"blog_gin/config"
	"blog_gin/core"
	"blog_gin/global"
	"blog_gin/pkg/constant/res_const"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) GetJwt(c *gin.Context) {
	jwt := global.Config.Jwt
	res.OkWithData(c, jwt)
}

func (SettingsApi) UpdateJwt(c *gin.Context) {
	var cr config.Jwt
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.Fail(c, res_const.INVALID_PARAMS, "")
		return
	}
	// 只将这一部分修改
	global.Config.Jwt = cr
	if !core.SetYaml() {
		res.Fail(c, res_const.FAIL_OPER, "")
		return
	}
	res.Ok(c)
}
