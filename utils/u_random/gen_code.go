package u_random

import (
	"math/rand"
	"time"
)

// GenRandomCode 生成指定长度的随机验证码
func GenRandomCode(length int) string {
	// 定义可能出现在验证码中的字符集
	charset := "0123456789"
	code := make([]byte, length)
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		code[i] = charset[randomIndex]
	}
	return string(code)
}
