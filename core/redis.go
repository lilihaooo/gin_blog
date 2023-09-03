package core

import (
	"blog_gin/global"
	"context"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	// 初始化 Redis 连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "a111111",        // Redis 密码
		DB:       0,                // 使用的数据库索引
		//ReadTimeout:  20 * time.Second, // 读取超时时间
		//WriteTimeout: 20 * time.Second,
	})
	ctx := context.Background()
	// 确保连接正常
	_, err := client.Ping(ctx).Result()
	if err != nil {
		global.Logrus.Fatal("redis连接失败:", err.Error())
	}
	global.RedisClient = client
}
