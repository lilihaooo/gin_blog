package menu_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"github.com/gin-gonic/gin"
)

type bannerSort struct {
	BannerID uint `json:"banner_id" validate:"required" label:"图片ID"`
	Sort     int  `json:"sort" validate:"required" label:"图片排序"`
}
type menuCreateRequest struct {
	Title          string       `validate:"required,max=30"  label:"标题"`
	Path           string       `validate:"required"` // 跳转路径
	Slogan         string       // 标语
	Abstract       ctype.Array  // 摘要
	AbstractTime   int          `validate:"required"` // 切换时间
	BannerTime     int          // 切换时间
	Sort           int          `validate:"required"`
	BannerSortList []bannerSort `validate:"dive"` //进入嵌套结构体中验证 // 带排序的bannerID 列表
}

func (MenuApi) MenuCreate(c *gin.Context) {
	var cr menuCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}

	// 重复值判断
	err := global.DB.Select("title").Where("title = ?", cr.Title).Take(&models.MenuModel{}).Error
	if err == nil {
		res.Fail(c, res.INVALID_PARAMS, "标题已存在")
		return
	}
	err = global.DB.Select("path").Where("path = ?", cr.Path).Take(&models.MenuModel{}).Error
	if err == nil {
		res.Fail(c, res.INVALID_PARAMS, "跳转路径已存在")
		return
	}

	// 添加Menu数据入库
	menu := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	if err = global.DB.Create(&menu).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "菜单添加失败")
		return
	}

	if len(cr.BannerSortList) == 0 {
		res.OkWithMsg(c, "添加成功")
		return
	}
	
	var menuBanners []models.MenuBannerModel
	// 将图片上传的banner_ids和menu关联起来保存到关系表menu_banner关系表中
	for _, item := range cr.BannerSortList {
		menuBanners = append(menuBanners, models.MenuBannerModel{
			MenuID:   menu.ID,
			BannerID: item.BannerID,
			Sort:     item.Sort,
		})
	}
	if err = global.DB.Create(&menuBanners).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "菜单图片关联失败")
		return
	}
	res.OkWithMsg(c, "添加成功")
}
