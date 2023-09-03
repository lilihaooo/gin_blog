package user_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/req"
	"blog_gin/pkg/res"
	"blog_gin/service/common"
	"blog_gin/utils/desensitize"
	"blog_gin/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserList(c *gin.Context) {
	// 获取用户角色
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)

	cr := req.NewPaginationReq()
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	var model models.UserModel
	sort := "id desc"
	db := global.DB.Order(sort)
	list, count, err := common.MakeList(model, db, cr)
	if err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "查询失败")
		return
	}
	for i, item := range list {
		if payload.Role == int(ctype.PermissionUser) {
			list[i].UserName = ""
		}
		// 电话脱敏
		list[i].Tel = desensitize.DesensitizeTel(item.Tel)
		// 邮箱脱敏
		list[i].Email = desensitize.DesensitizeEmail(item.Email)
	}

	res.OkWithList(c, list, count)
}
