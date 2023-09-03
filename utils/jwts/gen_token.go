package jwts

import (
	"blog_gin/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// GenToken 创建token
func GenToken(user Payload) string {
	jwtSecret := []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		Payload: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(GetJwtExpiresDuration())), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 生成令牌
	tokenString, err := token.SignedString(jwtSecret)         // 将令牌签名
	if err != nil {
		global.Logrus.Error("创建token失败:", err)
		return ""
	}
	return tokenString
}
