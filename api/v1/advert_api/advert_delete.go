package advert_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/req"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (AdvertsApi) AdvertDelete(c *gin.Context) {
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
	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDs).RowsAffected
	if count == 0 {
		res.Fail(c, res.INVALID_PARAMS, "广告不存在")
		return
	}
	if err := global.DB.Delete(advertList).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "删除失败")
		return
	}
	res.OkWithMsg(c, fmt.Sprintf("共删除%d条数据", count))
}
