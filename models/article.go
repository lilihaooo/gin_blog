package models

import (
	"blog_gin/global"
	"blog_gin/models/ctype"
	"blog_gin/pkg/req"
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID        string `json:"id"`         // es的id
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间

	Title    string `json:"title"`              // 文章标题
	Keyword  string `json:"keyword"`            // 文章标题
	Abstract string `json:"abstract"`           // 文章简介
	Content  string `json:"content,omit(list)"` // 文章内容

	LookCount     int `json:"look_count"`     // 浏览量
	CommentCount  int `json:"comment_count"`  // 评论量
	DiggCount     int `json:"digg_count"`     // 点赞量
	CollectsCount int `json:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id"`        // 用户id
	UserNickName string `json:"user_nick_name"` //用户昵称
	UserAvatar   string `json:"user_avatar"`    // 用户头像

	Category string `json:"category"`          // 文章分类
	Source   string `json:"source,omit(list)"` // 文章来源
	Link     string `json:"link,omit(list)"`   // 原文链接

	BannerID  uint   `json:"banner_id"`  // 文章封面id
	BannerUrl string `json:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags"` // 文章标签
}

func (a ArticleModel) TableName() string {
	// 自定义表名的逻辑
	return "article"
}

func (a ArticleModel) Index() string {
	return "article_index"
}

func (a ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "keyword"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "keyword"
      },
	  "tags": { 
		"type": "keyword"
	  },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (a ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		a.RemoveIndex()
	}
	// 没有索引
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil
}

// RemoveIndex 删除索引
func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}

// Create 添加的方法
func (a ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(a.Index()).
		BodyJson(a).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (a ArticleModel) ISExistData() bool {
	res, err := global.ESClient.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (a ArticleModel) EsGetArticleDetailByTitle() (model ArticleModel, err error) {
	searchResult, err := global.ESClient.Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if searchResult.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := searchResult.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	return model, nil
}

func (a ArticleModel) EsGetArticleDetail(id string) (model ArticleModel, err error) {
	searchResult, err := global.ESClient.Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(searchResult.Source, &model)
	if err != nil {
		return
	}
	model.ID = searchResult.Id
	return model, nil
}
func (a ArticleModel) EsGetArticleList(pageReq *req.PaginationReq, key string) (list []ArticleModel, count int64, err error) {
	offset := req.GetOffset(pageReq)
	if offset < 0 {
		offset = 0
	}
	boolQuery := elastic.NewBoolQuery()
	if key != "" {
		boolQuery.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	searchResult, err := global.ESClient.Search().
		Index(a.Index()).
		Query(boolQuery). // 查询条件
		From(offset).Size(pageReq.PageSize). // 分页
		Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	count = searchResult.Hits.TotalHits.Value
	if count == 0 {
		return []ArticleModel{}, count, nil
	}
	for _, hit := range searchResult.Hits.Hits {
		err = json.Unmarshal(hit.Source, &a)
		if err != nil {
			return nil, 0, err
		}
		a.ID = hit.Id
		list = append(list, a)
	}
	return list, count, nil

}
