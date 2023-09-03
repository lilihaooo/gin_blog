package advert_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"github.com/gin-gonic/gin"
)

type AdvertCreateRequest struct {
	Title  string `json:"title" validate:"required,max=20" label:"标题"` // 显示的标题
	Href   string `json:"href" validate:"required,url" label:"跳转链接"`   // 跳转链接
	Images string `json:"images" validate:"required,url" label:"图片地址"` // 图片地址
	IsShow bool   `json:"is_show"  label:"是否展示"`                       // 是否展示
}

func (AdvertsApi) AdvertCreate(c *gin.Context) {
	var cr AdvertCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	// 判断标题是否存在
	err := global.DB.Select("title").Where("title = ?", cr.Title).Take(&models.AdvertModel{}).Error
	if err == nil {
		res.Fail(c, res.INVALID_PARAMS, "广告已存在")
		return
	}

	advert := models.AdvertModel{
		Href:   cr.Href,
		Title:  cr.Title,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}

	if err = global.DB.Create(&advert).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "添加失败")
		return
	}
	res.OkWithMsg(c, "添加成功")
}
