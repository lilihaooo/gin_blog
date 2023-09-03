package menu_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetail(c *gin.Context) {
	menuId := c.Param("id")
	if menuId == "" {
		res.Fail(c, res.INVALID_PARAMS, "id不能为空")
		return
	}

	var menu models.MenuModel
	if err := global.DB.Take(&menu, menuId).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "记录不存在")
		return
	}

	var menuBanners []models.MenuBannerModel
	if err := global.DB.Preload("BannerModel").Order("sort desc").Where("menu_id = ?", menuId).Find(&menuBanners).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "查询失败")
		return
	}

	banners := []Banner{}
	for _, menuBanner := range menuBanners {
		if menu.ID == menuBanner.MenuID {
			banner := Banner{
				BannerID: menuBanner.BannerID,
				Path:     menuBanner.BannerModel.Path,
			}
			banners = append(banners, banner)
		}

	}
	response := menuResponse{
		ID:           menu.ID,
		Slogan:       menu.Slogan,
		Abstract:     menu.Abstract,
		AbstractTime: menu.AbstractTime,
		BannerTime:   menu.BannerTime,
		Sort:         menu.Sort,
		BannerList:   banners,
	}
	res.OkWithData(c, response)
}
