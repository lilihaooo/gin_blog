package config

// Config 配置文件
type Config struct {
	Mysql    `yaml:"mysql"`
	Logger   `yaml:"logger"`
	Server   `yaml:"server"`
	SiteInfo `yaml:"site_info"`
	QQ       `yaml:"qq"`
	Qiniu    `yaml:"qiniu"`
	Email    `yaml:"email"`
	Jwt      `yaml:"jwt"`
}

// ErrMap 错误映射
type ErrMap map[int]string
