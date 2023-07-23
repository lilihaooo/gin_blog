package core

import (
	"blog_gin/config"
	"blog_gin/global"
	"encoding/json"
	"os"
)

func InitErrorMap() {
	errMap := config.ErrMap{}
	// 读取错误码json文件到global中
	errorFlie := "error_code.json"
	byteData, err := os.ReadFile(errorFlie)
	if err != nil {
		global.Logrus.Fatal("错误码json文件读取失败")
	}
	// 将数据写入map中
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		global.Logrus.Fatal("解析错误文件失败")
	}
	global.ErrMap = &errMap
}
