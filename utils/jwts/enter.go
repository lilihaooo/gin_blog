package jwts

import (
	"blog_gin/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type Payload struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
}

type CustomClaims struct {
	Payload
	jwt.StandardClaims
}

// GetJwtExpiresDuration 获得jwt过期时间就间隔
func GetJwtExpiresDuration() time.Duration {
	return time.Minute * time.Duration(global.Config.Jwt.Expires)
}
