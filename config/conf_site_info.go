package config

type SiteInfo struct {
	CreatedAt string `yaml:"created_at" json:"created_at"`
	BeiAn     string `yaml:"bei_an" json:"bei_an"`
	Title     string `yaml:"title" json:"title"`
	Addr      string `yaml:"addr" json:"addr"`
}
