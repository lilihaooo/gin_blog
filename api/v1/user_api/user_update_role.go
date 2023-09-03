package user_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"blog_gin/utils/jwts"
	"blog_gin/utils/u_redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateRoleRequest struct {
	ID       uint       `json:"id" validate:"required"`
	Role     ctype.Role `validate:"required,oneof=1 2 3 4"`
	NickName string     `json:"nick_name" validate:"omitempty,max=36"` // 当管理员认为昵称不合法时会修改其昵称
}

func (UserApi) UserUpdateRole(c *gin.Context) {
	// 获取用户角色
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	if ctype.Role(payload.Role) != ctype.PermissionAdmin {
		res.Fail(c, res.ERROR_AUTH_CHECK_FAIL, "")
		return
	}
	var cr updateRoleRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 1号用户永远都是超级管理员
	if cr.ID == 1 {
		res.Fail(c, res.FAIL_OPER, "超级管理员, 不允许修改")
		return
	}

	vRes := utils.ZhValidate(&cr)
	if vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	// 查询是否存在
	var model models.UserModel
	err := global.DB.Take(&model, cr.ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res.Fail(c, res.ERROR_NOT_EXIST_USER, "")
			return
		}
		global.Logrus.Error(err)
		return
	}

	data := map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}
	if err = global.DB.Model(&model).Updates(data).Error; err != nil {
		res.Fail(c, res.FAIL_OPER, "修改角色失败")
		return
	}

	// 删除该用户所有token
	matchKey := fmt.Sprintf("jwt_token:%d:*", model.ID)
	if model.Role == ctype.PermissionDisableUser {
		u_redis.DeleteAllKeys(global.RedisClient, matchKey)
	}
	res.Ok(c)
}
