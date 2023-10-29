package global

import (
	"blog_gin/config"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config      *config.Config
	DB          *gorm.DB
	Logrus      *logrus.Logger
	ResMap      *config.ResMap
	RedisClient *redis.Client
	ESClient    *elastic.Client
)
