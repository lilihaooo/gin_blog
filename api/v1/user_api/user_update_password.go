package user_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"blog_gin/utils/jwts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	// 获取用户角色
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	// 鉴权
	if ctype.Role(payload.Role) != ctype.PermissionAdmin && ctype.Role(payload.Role) != ctype.PermissionUser {
		res.Fail(c, res.ERROR_AUTH_CHECK_FAIL, "")
		return
	}
	var cr updatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	vRes := utils.ZhValidate(&cr)
	if vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	if cr.Password == cr.OldPassword {
		res.Fail(c, res.INVALID_PARAMS, "两次输入的密码一致")
		return
	}

	// 查询user是否存在
	var model models.UserModel
	err := global.DB.Take(&model, payload.UserID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res.Fail(c, res.ERROR_NOT_EXIST_USER, "")
			return
		}
		global.Logrus.Error(err)
		return
	}
	// 对旧密码进行验证
	if ok := utils.CheckPasswordHash(cr.OldPassword, model.Password); !ok {
		res.Fail(c, res.FAIL_OPER, "密码错误!!!")
		return
	}
	// 对新密码进行加密处理
	hashPassword, err := utils.HashPassword(cr.Password)
	if err != nil {
		res.Fail(c, res.FAIL_OPER, "密码加密失败!!!")
		return
	}
	if err = global.DB.Model(&model).Update("password", hashPassword).Error; err != nil {
		res.Fail(c, res.FAIL_OPER, "修改密码失败")
		return
	}
	res.Ok(c)
}
