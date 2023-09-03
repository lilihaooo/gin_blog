package user_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"blog_gin/utils/jwts"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	UserName string `json:"user_name" validate:"max=36" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

func (UserApi) Login(c *gin.Context) {
	var cr LoginRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	// 查找用户是否为空
	var user models.UserModel
	if err := global.DB.Take(&user, "user_name = ?", cr.UserName).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			res.Fail(c, res.ERROR_NOT_EXIST_USER, "")
			return
		}
		global.Logrus.Error(err)
		return
	}
	// 验证密码是否正确
	ok := utils.CheckPasswordHash(cr.Password, user.Password)
	if !ok {
		res.Fail(c, res.ERROR_PASS_USER, "")
		return
	}
	payload := jwts.Payload{
		UserID:   user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Role:     int(user.Role),
	}
	token := jwts.GenToken(payload)
	key := fmt.Sprintf("jwt_token:%d:%s", user.ID, token)
	// 将jwt保存到redis中
	global.RedisClient.Set(context.Background(), key, "", jwts.GetJwtExpiresDuration())
	response := map[string]any{
		"token": token,
	}
	res.OkWithData(c, response)

}
