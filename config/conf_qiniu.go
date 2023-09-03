package config

type QiNiu struct {
	Enable    bool    `json:"enable" yaml:"enable"` // 是否启用七牛云储存
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` // 储存桶的名字
	CDN       string  `json:"cdn" yaml:"cdn"`       // 访问图片的地址前缀
	Zone      string  `json:"zone" yaml:"zone"`     // 储存地区
	Size      float64 `json:"size" yaml:"size"`     // 储存的大小限制, 单位是MB
	Prefix    string  `json:"prefix" yaml:"prefix"` // 储存的大小限制, 单位是MB
}
