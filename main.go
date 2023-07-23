package main

import (
	"blog_gin/core"
	"blog_gin/global"
	"blog_gin/routers"
)

func main() {
	core.InitConf()
	errorFile := core.InitLogrus()
	defer errorFile.Close()
	gormLogFile := core.InitGorm()
	defer gormLogFile.Close()
	// 初始化错误码json文件
	core.InitErrorMap()
	//fmt.Println((*global.ErrMap)[error_const.SUCCESS])

	r := routers.InitRouter()
	addr := global.Config.Server.Addr()
	global.Logrus.Infof("http server 运行在: %s", addr)
	r.Run(addr)
}
