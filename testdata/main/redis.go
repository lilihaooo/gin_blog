package main

import (
	"blog_gin/core"
	"blog_gin/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main3() {
	// 初始化配置文件
	core.InitConf()
	// 初始化日志
	core.InitLogrus()
	defer core.CloseLogFile()
	// 初始化Gorm
	core.InitGorm()
	defer core.CloseGormLogFile()
	// 初始化错误码json文件
	core.InitResMap()
	//// 初始化redis
	core.InitRedis()
	//id := 1
	//for {
	//	token := "asasasda" + strconv.FormatInt(time.Now().Unix(), 10)
	//	key := fmt.Sprintf("user:%d:%s", id, token)
	//	result, err := global.RedisClient.Set(context.Background(), key, "", time.Hour).Result()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(result)
	//}
	keys, _ := getAllKeys(global.RedisClient, "user:1:*")
	fmt.Println(len(keys))
	for _, key := range keys {
		fmt.Println(key)
	}

}
func getAllKeys(client *redis.Client, matchKey string) ([]string, error) {
	var keys []string
	var cursor uint64 = 0
	ctx := context.Background()

	for {
		var scanKeys []string
		var err error
		scanKeys, cursor, err = client.Scan(ctx, cursor, matchKey, 10).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, scanKeys...)
		if cursor == 0 {
			break
		}
	}

	return keys, nil
}
