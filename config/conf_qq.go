package config

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` // 登陆之后回调的地址
}

func (q *QQ) GetPath() string {
	if q.Key == "" || q.AppID == " " || q.Redirect == "" {
		return ""
	}
	// todo 网址
	return fmt.Sprintf("https://qraph.qq.com/")
}
