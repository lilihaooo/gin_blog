package menu_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	BannerID uint   `json:"banner_id"`
	Path     string `json:"path"`
}
type menuResponse struct {
	ID           uint        `json:"id"`
	Title        string      `json:"title"`
	Slogan       string      `json:"slogan"`
	Abstract     ctype.Array `json:"abstract"`
	AbstractTime int         `json:"abstract_time"`
	BannerTime   int         `json:"banner_time"`
	Sort         int         `json:"sort"`
	BannerList   []Banner    `json:"banner"`
}

func (MenuApi) MenuList(c *gin.Context) {
	//  查询menu list
	var menus []models.MenuModel
	var MenuIDs []uint
	if err := global.DB.Order("sort desc").Find(&menus).Select("id").Scan(&MenuIDs).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "查询失败")
		return
	}
	// 根据查询出所有的图片
	var menuBanners []models.MenuBannerModel
	if err := global.DB.Preload("BannerModel").Order("sort desc").Where("menu_id in ?", MenuIDs).Find(&menuBanners).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "查询失败")
		return
	}

	// 遍历两个表, 将数据对应起来
	response := []menuResponse{}
	for _, menu := range menus {
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
		response = append(response, menuResponse{
			ID:           menu.ID,
			Slogan:       menu.Slogan,
			Abstract:     menu.Abstract,
			AbstractTime: menu.AbstractTime,
			BannerTime:   menu.BannerTime,
			Sort:         menu.Sort,
			BannerList:   banners,
		})
	}
	res.OkWithList(c, response, int64(len(response)))
}
