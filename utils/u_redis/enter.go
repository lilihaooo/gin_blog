package u_redis

import (
	"blog_gin/global"
	"context"
	"github.com/go-redis/redis/v8"
)

func GetAllKeys(client *redis.Client, matchKey string) ([]string, error) {
	var keys []string
	var cursor uint64 = 0
	ctx := context.Background()
	for {
		scanKeys, newCursor, err := client.Scan(ctx, cursor, matchKey, 10).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, scanKeys...)
		if newCursor == 0 {
			break
		}
	}
	return keys, nil
}

func DeleteAllKeys(client *redis.Client, matchKey string) bool {
	keys, err := GetAllKeys(client, matchKey)
	if err != nil {
		global.Logrus.Error(err)
		return false
	}
	ctx := context.Background()
	for _, key := range keys {
		_, err := global.RedisClient.Del(ctx, key).Result()
		if err != nil {
			global.Logrus.Error(err)
			return false
		}
	}
	return true
}
