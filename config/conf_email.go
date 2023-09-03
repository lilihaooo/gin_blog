package config

type Email struct {
	From     string `json:"from" yaml:"from"` // 发件人邮箱
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
}
