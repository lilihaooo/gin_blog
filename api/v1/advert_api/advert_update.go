package advert_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type AdvertUpdateRequest struct {
	ID     int64  `json:"id" validate:"required"`                                       // ID
	Href   string `json:"href" validate:"required,url" label:"跳转链接" structs:"href"`     // 跳转链接
	Images string `json:"images" validate:"required,url" label:"图片地址" structs:"images"` // 图片地址
	IsShow bool   `json:"is_show" structs:"is_show"`                                    // 是否展示
}

func (AdvertsApi) AdvertUpdate(c *gin.Context) {
	var cr AdvertUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	// 判断是否存在
	var advert models.AdvertModel
	count := global.DB.Select("id").Take(&advert, cr.ID).RowsAffected
	if count != 1 {
		res.Fail(c, res.INVALID_PARAMS, "广告不存在")
		return
	}

	// 将结构体转为map
	crMap := structs.Map(&cr)
	// 要是传入的bool值, 例如 is_show 为false 使用结构体将会忽略掉该字段, 所以修改bool值应该使用map
	if err := global.DB.Model(advert).Updates(crMap).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "编辑失败")
		return
	}
	res.OkWithMsg(c, "编辑成功")
}
