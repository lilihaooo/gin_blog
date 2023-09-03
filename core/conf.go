package core

import (
	"blog_gin/config"
	"blog_gin/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var file = "settings.yaml"

// InitConf 读取yaml文件的配置
func InitConf() {
	// 读取配置文件
	file, err := os.Open(file)
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

// 写入yaml文件
func SetYaml() bool {
	// 将结构体保存回YAML文件
	newData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Logrus.Error("数据序列化失败", err)
		return false
	}
	if err = ioutil.WriteFile(file, newData, 0644); err != nil {
		return false
	}
	return true
}
