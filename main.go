package main

import (
	"blog_gin/core"
	"blog_gin/flag"
	"blog_gin/global"
	"blog_gin/routers"
)

func main() {
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

	option := flag.Parse()
	flag.SwitchOption(option)
	if flag.IsWebStop(option) {
		return
	}

	r := routers.InitRouter()
	addr := global.Config.Server.Addr()
	global.Logrus.Infof("http server 运行在: %s", addr)
	err := r.Run(addr)
	if err != nil {
		global.Logrus.Fatal(err.Error())
	}
}
