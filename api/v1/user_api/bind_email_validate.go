package user_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"blog_gin/utils/jwts"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BindEmailValidateRequest struct {
	Email string `json:"email" validate:"required,email" label:"邮箱"`
	Code  string `json:"code" validate:"required" label:"验证码"`
}

func (UserApi) BindEmailValidate(c *gin.Context) {
	var cr BindEmailValidateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	user := c.MustGet("user")
	payload := user.(*jwts.Payload)
	// 验证email
	hashName := "code:" + strconv.Itoa(int(payload.UserID))
	hashFields, err := global.RedisClient.HGetAll(c, hashName).Result()
	if err != nil {
		res.Fail(c, res.FAIL_OPER, "redis操作失败")
		return
	}
	if hashFields["code"] == "" {
		res.Fail(c, res.FAIL_VALIDATE_CODE, "code已过期")
		return
	}

	if hashFields["code"] != cr.Code {
		res.Fail(c, res.FAIL_VALIDATE_CODE, "验证码不正确")
		return
	}
	if hashFields["email"] != cr.Email {
		res.Fail(c, res.FAIL_VALIDATE_CODE, "email不一致!!!")
		return
	}

	var userModel models.UserModel
	userModel.ID = payload.UserID
	newData := map[string]interface{}{
		"email": cr.Email,
	}
	if row := global.DB.Model(&userModel).Updates(newData).RowsAffected; row == 0 {
		res.Fail(c, res.FAIL_OPER, "数据更新失败")
		return
	}
	res.Ok(c)
}
