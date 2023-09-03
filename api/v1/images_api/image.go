package images_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/req"
	"blog_gin/pkg/res"
	"blog_gin/service/common"
	"blog_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImagesList(c *gin.Context) {
	pageReq := req.NewPaginationReq()
	err := c.ShouldBindQuery(&pageReq)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	sort := "id desc"
	db := global.DB.Order(sort)
	list, count, err := common.MakeList(models.BannerModel{}, db, pageReq)
	if err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "查询失败")
		return
	}
	res.OkWithList(c, list, count)
}

func (ImagesApi) ImagesDelete(c *gin.Context) {
	var cr req.DeleteRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}

	if len(cr.IDs) == 0 {
		res.Fail(c, res.INVALID_PARAMS, "请选择图片")
		return
	}

	var bannerList []models.BannerModel
	count := global.DB.Find(&bannerList, cr.IDs).RowsAffected
	if count == 0 {
		res.Fail(c, res.INVALID_PARAMS, "图片不存在")
		return
	}
	err = global.DB.Delete(bannerList).Error
	if err != nil {
		// 钩子函数执行失败返回err
		res.Fail(c, res.FAIL_OPER, err.Error())
		return
	}
	res.OkWithMsg(c, fmt.Sprintf("共删除%d条数据", count))
}

func (ImagesApi) ImageUpdate(c *gin.Context) {
	var cr struct {
		ID   int64  `json:"id" validate:"required"`
		Name string `json:"name" validate:"required,max=8" label:"名称"`
	}
	// 接收参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 验证字段
	vRes := utils.ZhValidate(cr)
	if vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	var banner models.BannerModel
	err = global.DB.Take(&banner, cr.ID).Error
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "图片不存在")
		return
	}
	banner.Name = cr.Name
	global.DB.Updates(&banner)
	res.OkWithMsg(c, "图片修改成功")
}
