package core

import (
	"blog_gin/config"
	"blog_gin/global"
	"encoding/json"
	"os"
)

func InitResMap() {
	resMap := config.ResMap{}
	// 读取错误码json文件到global中
	responseFlie := "response.json"
	byteData, err := os.ReadFile(responseFlie)
	if err != nil {
		global.Logrus.Fatal("返回码json文件读取失败")
	}
	// 将数据写入map中
	err = json.Unmarshal(byteData, &resMap)
	if err != nil {
		global.Logrus.Fatal("解析错误文件失败")
	}
	global.ResMap = &resMap
}
