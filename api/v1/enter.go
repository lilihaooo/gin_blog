package v1

import (
	"blog_gin/api/v1/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
