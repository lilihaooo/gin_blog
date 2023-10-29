package config

import "fmt"

type Es struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (e Es) Url() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}
