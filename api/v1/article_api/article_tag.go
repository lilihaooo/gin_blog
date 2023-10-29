package article_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/req"
	"blog_gin/pkg/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type articleTagResponse struct {
	TagName   string        `json:"tag_name"`
	Count     int           `json:"count"`
	List      []articleList `json:"list"`
	CreatedAt string        `json:"created_at"`
}

type articleList struct {
	ArticleID   string `json:"tag_name"`
	ArticleName string `json:"article_name"`
}

type tagBuckets struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key       string `json:"key"`
		DocCount  int    `json:"doc_count"`
		ArticleID struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"article_id"`
		ArticleKey struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"article_key"`
	} `json:"buckets"`
}

// ArticleCalendar es文章按时间聚合
func (ArticleApi) ArticleTag(c *gin.Context) {
	cr := req.NewPaginationReq()
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	// 由于分页的原因所以先查询总数
	countResult, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tag_count", elastic.NewCardinalityAggregation().Field("tags")).
		Do(context.Background())
	cTag, _ := countResult.Aggregations.Cardinality("tag_count")
	count := *cTag.Value

	// 创建一个聚合查询
	agg := elastic.NewTermsAggregation().Field("tags") // 根据标签字段进行聚合
	agg.SubAggregation("article_id", elastic.NewTermsAggregation().Field("_id"))
	agg.SubAggregation("article_key", elastic.NewTermsAggregation().Field("keyword"))

	// 分页子聚合
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(req.GetOffset(cr)).Size(cr.PageSize))

	searchResult, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("articles_with_tag", agg).
		Do(context.Background())
	if err != nil {
		res.Fail(c, res.FAIL_OPER, err.Error())
		return
	}

	var buckets tagBuckets
	json.Unmarshal(searchResult.Aggregations["articles_with_tag"], &buckets)

	var tagStringList []string
	response := []*articleTagResponse{}
	for _, item := range buckets.Buckets {
		_articleList := make([]articleList, len(item.ArticleID.Buckets))
		for i, v := range item.ArticleKey.Buckets {
			_articleList[i].ArticleName = v.Key
		}
		for i, v := range item.ArticleID.Buckets {
			_articleList[i].ArticleID = v.Key
		}
		_data := &articleTagResponse{
			TagName: item.Key,
			Count:   item.DocCount,
			List:    _articleList,
		}
		tagStringList = append(tagStringList, item.Key)
		response = append(response, _data)
	}

	// 查询出所有的标签
	var tags []models.TagModel
	err = global.DB.Where("title in ?", tagStringList).Find(&tags).Error
	if err != nil {
		res.Fail(c, res.FAIL_OPER, "数据库查询失败")
		return
	}
	createdAtMap := make(map[string]string)
	for _, item := range tags {
		createdAtMap[item.Title] = item.CreatedAt.Format("2006-01-02 15:04:05")
	}
	for _, item := range response {
		item.CreatedAt = createdAtMap[item.TagName]
	}
	res.OkWithList(c, response, int64(count))

}
