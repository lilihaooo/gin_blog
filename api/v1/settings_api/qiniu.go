package settings_api

import (
	"blog_gin/config"
	"blog_gin/core"
	"blog_gin/global"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) GetQiNiu(c *gin.Context) {
	qiniu := global.Config.QiNiu
	res.OkWithData(c, qiniu)
}

func (SettingsApi) UpdateQiNiu(c *gin.Context) {
	var cr config.QiNiu
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 只将这一部分修改
	global.Config.QiNiu = cr
	if !core.SetYaml() {
		res.Fail(c, res.FAIL_OPER, "")
		return
	}
	res.Ok(c)
}
