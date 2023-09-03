package jwts

import (
	"blog_gin/global"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken 验证token
func ParseToken(tokenString string) (*CustomClaims, error) {
	jwtSecret := []byte(global.Config.Jwt.Secret)
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		global.Logrus.Error("解析token失败:", err)
		return nil, err
	}
	if !token.Valid {
		err = errors.New("token验证失败")
		global.Logrus.Error("token验证失败:", err)
		return nil, err
	}
	return &claims, nil
}
