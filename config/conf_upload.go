package config

type Upload struct {
	Size      float64  `json:"size" yaml:"size"`
	Path      string   `json:"path" yaml:"path"`
	WhiteList []string `json:"white_list" yaml:"white_list"`
}
