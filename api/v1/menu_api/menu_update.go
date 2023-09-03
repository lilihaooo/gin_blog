package menu_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type updateBannerSort struct {
	BannerID uint `json:"banner_id" validate:"required" label:"图片ID"`
	Sort     int  `json:"sort" validate:"required" label:"图片排序"`
}
type menuUpdateRequest struct {
	ID             uint               `json:"id" structs:"id"`
	Title          string             `validate:"required,max=30" structs:"title"`
	Path           string             `validate:"required" structs:"path"` // 跳转路径
	Slogan         string             `structs:"slogan"`                   // 标语
	Abstract       ctype.Array        `structs:"abstract"`                 // 介绍
	AbstractTime   int                `structs:"abstract_time"`            // 介绍切换时间
	BannerTime     int                `structs:"banner_time"`              // 图片切换时间
	Sort           int                `validate:"required" structs:"sort"`
	BannerSortList []updateBannerSort `validate:"dive" structs:"-"` //进入嵌套结构体中验证 // 带排序的bannerID 列表
}

func (MenuApi) MenuUpdate(c *gin.Context) {
	var cr menuUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		fmt.Println(err.Error())
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}

	var menu models.MenuModel
	count := global.DB.Select("id").Take(&menu, cr.ID).RowsAffected
	if count != 1 {
		res.Fail(c, res.INVALID_PARAMS, "菜单不存在")
		return
	}
	// 将结构体转为map
	crMap := structs.Map(&cr)
	if err := global.DB.Model(menu).Updates(&crMap).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "菜单修改失败")
		return
	}

	// 删除关联表中该菜单的所有记录
	if err := global.DB.Model(&menu).Association("Banners").Clear(); err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "关联表删除失败")
		return
	}
	// 如果传入null,删除该菜单关联表中的记录
	if len(cr.BannerSortList) != 0 {
		// 将图片上传的banner_ids和menu关联起来保存到关系表menu_banner关系表中
		var menuBanners []models.MenuBannerModel
		for _, item := range cr.BannerSortList {
			menuBanners = append(menuBanners, models.MenuBannerModel{
				MenuID:   menu.ID,
				BannerID: item.BannerID,
				Sort:     item.Sort,
			})
		}
		if err := global.DB.Create(&menuBanners).Error; err != nil {
			global.Logrus.Error(err)
			res.Fail(c, res.FAIL_OPER, "菜单图片关联失败")
			return
		}
	}

	res.OkWithMsg(c, "修改成功")
}
