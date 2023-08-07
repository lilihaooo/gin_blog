package settings_api

import (
	"blog_gin/config"
	"blog_gin/core"
	"blog_gin/global"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) GetSittingsInfo(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(c, global.Config.SiteInfo)
	case "email":
		res.OkWithData(c, global.Config.Email)
	case "qq":
		res.OkWithData(c, global.Config.QQ)
	case "qiniu":
		res.OkWithData(c, global.Config.QiNiu)
	case "jwt":
		res.OkWithData(c, global.Config.Jwt)
	default:
		res.Fail(c, res.INVALID_PARAMS, "参数错误啦")
	}
}

func (SettingsApi) UpdateSittingsInfo(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}

	switch cr.Name {
	case "site":
		var data config.SiteInfo
		err = c.ShouldBindJSON(&data)
		if err != nil {
			res.Fail(c, res.INVALID_PARAMS, "")
			return
		}
		global.Config.SiteInfo = data
	case "email":
		var data config.Email
		err = c.ShouldBindJSON(&data)
		if err != nil {
			res.Fail(c, res.INVALID_PARAMS, "")
			return
		}
		global.Config.Email = data
	case "qq":
		var data config.QQ
		err = c.ShouldBindJSON(&data)
		if err != nil {
			res.Fail(c, res.INVALID_PARAMS, "")
			return
		}
		global.Config.QQ = data
	case "qiniu":
		var data config.QiNiu
		err = c.ShouldBindJSON(&data)
		if err != nil {
			res.Fail(c, res.INVALID_PARAMS, "")
			return
		}
		global.Config.QiNiu = data
	case "jwt":
		var data config.Jwt
		err = c.ShouldBindJSON(&data)
		if err != nil {
			res.Fail(c, res.INVALID_PARAMS, "")
			return
		}
		global.Config.Jwt = data
	default:
		res.Fail(c, res.INVALID_PARAMS, "参数错误啦")
		return
	}

	if !core.SetYaml() {
		res.Fail(c, res.FAIL_OPER, "")
		return
	}
	res.Ok(c)
}
