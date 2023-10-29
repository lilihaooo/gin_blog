package redis_ser

import (
	"blog_gin/global"
	"context"
)

// SetHashWithExpireTime 利用Lua脚本创建一个有过期时间的hash
func (RedisService) SetHashWithExpireTime(hashName string, args []interface{}) error {
	keys := []string{hashName} // 转换类型
	// 定义Lua脚本
	luaScript := `
        local key = KEYS[1]
        local expire = tonumber(ARGV[1])

        for i = 2, #ARGV, 2 do
            local field = ARGV[i]
            local value = ARGV[i + 1]
            redis.call('HSET', key, field, value)
        end

        if expire > 0 then
            redis.call('EXPIRE', key, expire)
        end
		return "" 
    `

	// 创建上下文
	ctx := context.Background()

	// 执行Lua脚本
	_, err := global.RedisClient.Eval(ctx, luaScript, keys, args).Result()
	if err != nil {
		return err
	}
	return nil
}
