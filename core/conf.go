package core

import (
	"blog_gin/config"
	"blog_gin/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	// 读取配置文件
	configFile := "settings.yaml"
	file, err := os.Open(configFile)
	if err != nil {
		panic(fmt.Errorf("无法打开配置文件: %s", err))

	}
	defer file.Close()
	// 解析配置文件
	var c config.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	global.Config = &c
}
