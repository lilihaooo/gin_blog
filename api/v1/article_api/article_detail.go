package article_api

import (
	"blog_gin/models"
	"blog_gin/pkg/res"
	"github.com/gin-gonic/gin"
)

// ArticleDetail 文章详情
func (ArticleApi) ArticleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		res.Fail(c, res.INVALID_PARAMS, "id不能为空")
		return
	}
	var model models.ArticleModel
	data, err := model.EsGetArticleDetail(id)
	if err != nil {
		res.Fail(c, res.FAIL_OPER, err.Error())
		return
	}
	res.OkWithData(c, data)
}

func (ArticleApi) ArticleDetailByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		res.Fail(c, res.INVALID_PARAMS, "title不能为空")
		return
	}
	var model models.ArticleModel
	model.Title = title
	data, err := model.EsGetArticleDetailByTitle()
	if err != nil {
		res.Fail(c, res.FAIL_OPER, err.Error())
		return
	}
	res.OkWithData(c, data)
}
