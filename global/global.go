package global

import (
	"blog_gin/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Logrus *logrus.Logger
	ResMap *config.ResMap
)
