package user_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email string `json:"email" validate:"required,email" label:"邮箱"`
}

func (UserApi) SendCode(c *gin.Context) {
	var cr BindEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	var user models.UserModel
	if row := global.DB.Where("email = ?", cr.Email).Find(user).RowsAffected; row > 0 {
		res.Fail(c, res.INVALID_PARAMS, "该email已被绑定")
	}
}
