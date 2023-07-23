package config

// todo 有没有用的
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	Show         string `yaml:"show"`
	ShowLine     string `yaml:"line"`
	LogInConsole bool   `yaml:"console"`
}
