package config

// Config 配置文件
type Config struct {
	Mysql  `yaml:"mysql"`
	Logger `yaml:"logger"`
	Server `yaml:"server"`
}

// ErrMap 错误映射
type ErrMap map[int]string
