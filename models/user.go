package models

import (
	"blog_gin/models/ctype"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	NickName   string           `gorm:"size:36;comment:昵称" json:"nick_name"`                          // 昵称
	UserName   string           `gorm:"size:36;comment:用户名" json:"user_name"`                         // 用户名
	Password   string           `gorm:"size:128;comment:密码" json:"password"`                          // 密码
	Avatar     string           `gorm:"size:256;comment:头像id" json:"avatar_id"`                       // 头像id
	Email      string           `gorm:"size:128;comment:邮箱" json:"email"`                             // 邮箱
	Tel        string           `gorm:"size:18;comment:手机号" json:"tel"`                               // 手机号
	Addr       string           `gorm:"size:64;comment:地址" json:"addr"`                               // 地址
	Token      string           `gorm:"size:64;comment:其他平台的唯一id" json:"token"`                       // 其他平台的唯一id
	IP         string           `gorm:"size:20;comment:ip地址" json:"ip"`                               // ip地址
	Role       ctype.Role       `gorm:"size:4;default:1;comment:权限  1 管理员  2 普通用户  3 游客" json:"role"` // 权限  1 管理员  2 普通用户  3 游客
	SignStatus ctype.SignStatus `gorm:"type=smallint(6);comment:登陆状态" json:"sign_status"`             // 登陆状态
}

func (m *UserModel) TableName() string {
	// 自定义表名的逻辑
	return "user"
}
