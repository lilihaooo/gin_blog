package advert_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/req"
	"blog_gin/pkg/res"
	"blog_gin/service/common/db_ser"
	"github.com/gin-gonic/gin"
	"strings"
)

func (AdvertsApi) AdvertList(c *gin.Context) {
	cr := req.NewPaginationReq()
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 如果是后端请求 显示全部的广告
	referer := c.GetHeader("Referer")
	isShow := true // 为true的话就要筛选该字段, == where("is_show = true")  为false 是会忽略该字段的where条件
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	model := models.AdvertModel{IsShow: isShow}
	sort := "id desc"
	db := global.DB.Order(sort).Where(model)
	list, count, err := db_ser.DBMakeList(model, db, cr)
	if err != nil {
		global.Logrus.Error(err)
		res.Fail(c, res.FAIL_OPER, "查询失败")
		return
	}
	res.OkWithList(c, list, count)
}
