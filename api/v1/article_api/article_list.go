package article_api

import (
	"blog_gin/models"
	"blog_gin/pkg/req"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

// ArticleList 文章列表
func (ArticleApi) ArticleList(c *gin.Context) {
	cr := req.NewPaginationReq()
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	key := c.Query("key")
	var model models.ArticleModel
	list, count, err := model.EsGetArticleList(cr, key)
	if err != nil {
		res.Fail(c, res.FAIL_OPER, err.Error())
		return
	}
	res.OkWithList(c, list, count)
}
