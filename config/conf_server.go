package config

import "fmt"

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
