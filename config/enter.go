package config

type config struct {
	Mysql  `yaml:"mysql"`
	Logger `yaml:"logger"`
	Server `yaml:"server"`
}
