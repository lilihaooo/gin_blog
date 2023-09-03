package menu_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

type menuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (MenuApi) MenuNameList(c *gin.Context) {
	var response []menuNameResponse
	if err := global.DB.Model(models.MenuModel{}).Order("sort desc").Select("id", "title", "path", "sort").Scan(&response).Error; err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "查询失败")
		return
	}
	res.OkWithData(c, response)
}
