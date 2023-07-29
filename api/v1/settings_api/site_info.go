package settings_api

import (
	"blog_gin/config"
	"blog_gin/core"
	"blog_gin/global"
	"blog_gin/pkg/constant/res_const"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) GetSiteInfo(c *gin.Context) {
	siteInfo := global.Config.SiteInfo
	res.OkWithData(c, siteInfo)
}

func (SettingsApi) UpdateSiteInfo(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.Fail(c, res_const.INVALID_PARAMS, "")
		return
	}
	// 只将这一部分修改
	global.Config.SiteInfo = cr
	if !core.SetYaml() {
		res.Fail(c, res_const.FAIL_OPER, "")
		return
	}
	res.Ok(c)
}
