package user_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"blog_gin/service"
	"blog_gin/utils"
	"blog_gin/utils/jwts"
	"blog_gin/utils/u_email"
	"blog_gin/utils/u_random"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BindEmailSendRequest struct {
	Email string `json:"email" validate:"required,email" label:"邮箱"`
}

func (UserApi) BindEmailSend(c *gin.Context) {
	var cr BindEmailSendRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	var userModel models.UserModel
	if row := global.DB.Where("email = ?", cr.Email).Find(&userModel).RowsAffected; row > 0 {
		res.Fail(c, res.INVALID_PARAMS, "该email已被绑定")
		return
	}

	//发送验证码
	code := u_random.GenRandomCode(4)
	if err := u_email.SendEmail("验证码", code, cr.Email); err != nil {
		res.Fail(c, res.FAIL_SEND_EMAIl, "")
		return
	}
	user := c.MustGet("user")
	payload := user.(*jwts.Payload)

	// 将id为key, code和email为字段以hash类型保存到redis中,同时过期时间为60s 利用lua脚本保证添加与设置过期时间的原子性
	hashName := "code:" + strconv.FormatUint(uint64(payload.UserID), 10) // 将 uint 转换为 string
	// lua脚本的参数格式{过期时间(s), key1, value1, key2, value2, ...}
	args := []interface{}{6000, "code", code, "email", cr.Email} // 参数列表
	err := service.AppService.SetHashWithExpireTime(hashName, args)
	if err != nil {
		res.Fail(c, res.FAIL_OPER, "redis操作失败"+err.Error())
		return
	}
	res.Ok(c)
}
