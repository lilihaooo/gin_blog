package main

import (
	"blog_gin/core"
	"blog_gin/utils/jwts"
	"fmt"
	"time"
)

func main1() {
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
	// 初始化redis
	core.InitRedis()

	user := jwts.Payload{
		UserID:   222,
		UserName: "test",
		NickName: "test",
		Role:     4,
	}
	token := jwts.GenToken(user)
	fmt.Println(token)
	time.Sleep(time.Second * 4)
	payload, err := jwts.ParseToken(token)
	fmt.Println(payload)
	fmt.Println(err)
}
