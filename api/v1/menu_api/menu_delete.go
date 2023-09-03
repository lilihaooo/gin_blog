package menu_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/req"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDelete(c *gin.Context) {
	var cr req.DeleteRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 验证参数
	if vRes := utils.ZhValidate(cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDs).RowsAffected
	if count == 0 {
		res.Fail(c, res.INVALID_PARAMS, "不存在")
		return
	}
	tx := global.DB.Begin()
	if err := tx.Model(menuList).Association("Banners").Clear(); err != nil {
		tx.Rollback()
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "关联表删除失败")
		return
	}
	result := tx.Delete(menuList)
	if err := result.Error; err != nil {
		tx.Rollback()
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "删除失败")
		return
	}
	rowsAffected := result.RowsAffected
	tx.Commit()
	res.OkWithMsg(c, fmt.Sprintf("共删除%d条数据", rowsAffected))
}
