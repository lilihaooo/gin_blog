package user_api

import (
	"blog_gin/global"
	"blog_gin/pkg/res"
	"blog_gin/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) Logout(c *gin.Context) {
	token := c.GetHeader("token")
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	key := fmt.Sprintf("jwt_token:%d:%s", payload.UserID, token)
	_, err := global.RedisClient.Del(c, key).Result()
	if err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "注销失败")
		return
	}
	res.Ok(c)
}
