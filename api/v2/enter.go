package v2

import (
	"blog_gin/api/v2/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
