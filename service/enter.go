package service

import (
	"blog_gin/service/image_ser"
	"blog_gin/service/redis_ser"
	"blog_gin/service/websocket_ser"
)

type Service struct {
	image_ser.ImageService
	websocket_ser.WebsocketService
	redis_ser.RedisService
}

var AppService = new(Service)
