package settings_api

import (
	"blog_gin/config"
	"blog_gin/core"
	"blog_gin/global"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) GetQQ(c *gin.Context) {
	qq := global.Config.QQ
	res.OkWithData(c, qq)
}

func (SettingsApi) UpdateQQ(c *gin.Context) {
	var cr config.QQ
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 只将这一部分修改
	global.Config.QQ = cr
	if !core.SetYaml() {
		res.Fail(c, res.FAIL_OPER, "")
		return
	}
	res.Ok(c)
}
