package main

import (
	"blog_gin/core"
	"blog_gin/utils/u_email"
	"blog_gin/utils/u_random"
	"fmt"
)

func main4() {
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

	if err := u_email.SendEmail(string(u_email.Code), u_random.GenRandomCode(4), "1039007652@qq.com"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("发送成功")
}
