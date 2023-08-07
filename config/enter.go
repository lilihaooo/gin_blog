package config

// Config 配置文件
type Config struct {
	Mysql    `yaml:"mysql"`
	Logger   `yaml:"logger"`
	Server   `yaml:"server"`
	SiteInfo `yaml:"site_info"`
	QQ       `yaml:"qq"`
	QiNiu    `yaml:"qi_niu"`
	Email    `yaml:"email"`
	Jwt      `yaml:"jwt"`
	Upload   `yaml:"upload"`
}

// ResMap 错误映射
type ResMap map[int]string
