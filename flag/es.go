package flag

import (
	"blog_gin/models"
)

func CreateEsIndex() {
	var articleModle models.ArticleModel
	err := articleModle.CreateIndex()
	if err != nil {
		panic("创建article索引失败")
	}
}

func DeleteEsIndex() {
	var articleModle models.ArticleModel
	err := articleModle.RemoveIndex()
	if err != nil {
		panic("移除article索引失败")
	}
}
