package models

import (
	"blog_gin/models/ctype"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model
	Title    string `gorm:"comment:文章标题" json:"title" structs:"title"`       // 文章标题
	Keyword  string `gorm:"comment:关键字" json:"keyword" structs:"keyword"`    // 关键字
	Abstract string `gorm:"comment:文章简介" json:"abstract" structs:"abstract"` // 文章简介
	Content  string `gorm:"comment:文章内容" json:"content" structs:"content"`   // 文章内容

	LookCount     int `gorm:"comment:浏览量" json:"look_count" structs:"look_count"`         // 浏览量
	CommentCount  int `gorm:"comment:评论量" json:"comment_count" structs:"comment_count"`   // 评论量
	DiggCount     int `gorm:"comment:点赞量" json:"digg_count" structs:"digg_count"`         // 点赞量
	CollectsCount int `gorm:"comment:收藏量" json:"collects_count" structs:"collects_count"` // 收藏量

	UserID       uint   `gorm:"comment:用户id" json:"user_id" structs:"user_id"`               // 用户id
	UserNickName string `gorm:"comment:用户昵称" json:"user_nick_name" structs:"user_nick_name"` //用户昵称
	UserAvatar   string `gorm:"comment:用户头像" json:"user_avatar" structs:"user_avatar"`       // 用户头像

	Category string `gorm:"comment:文章分类" json:"category" structs:"category"` // 文章分类
	Source   string `gorm:"comment:文章来源" json:"source" structs:"source"`     // 文章来源
	Link     string `gorm:"comment:原文链接" json:"link" structs:"link"`         // 原文链接

	BannerID  uint   `gorm:"comment:文章封面id" json:"banner_id" structs:"banner_id"` // 文章封面id
	BannerUrl string `gorm:"comment:文章封面" json:"banner_url" structs:"banner_url"` // 文章封面

	Tags ctype.Array `gorm:"comment:文章标签" json:"tags" structs:"tags"` // 文章标签
}

func (m *ArticleModel) TableName() string {
	// 自定义表名的逻辑
	return "article"
}

//
//func (ArticleModel) Index() string {
//	return "article_index"
//}
//
//func (ArticleModel) Mapping() string {
//	return `
//{
//  "settings": {
//    "index":{
//      "max_result_window": "100000"
//    }
//  },
//  "mappings": {
//    "properties": {
//      "title": {
//        "type": "text"
//      },
//      "keyword": {
//        "type": "keyword"
//      },
//      "abstract": {
//        "type": "text"
//      },
//      "content": {
//        "type": "text"
//      },
//      "look_count": {
//        "type": "integer"
//      },
//      "comment_count": {
//        "type": "integer"
//      },
//      "digg_count": {
//        "type": "integer"
//      },
//      "collects_count": {
//        "type": "integer"
//      },
//      "user_id": {
//        "type": "integer"
//      },
//      "user_nick_name": {
//        "type": "keyword"
//      },
//      "user_avatar": {
//        "type": "keyword"
//      },
//      "category": {
//        "type": "keyword"
//      },
//      "source": {
//        "type": "keyword"
//      },
//      "link": {
//        "type": "keyword"
//      },
//      "banner_id": {
//        "type": "integer"
//      },
//      "banner_url": {
//        "type": "keyword"
//      },
//      "tags": {
//        "type": "keyword"
//      },
//      "created_at":{
//        "type": "date",
//        "null_value": "null",
//        "format": "[yyyy-MM-dd HH:mm:ss]"
//      },
//      "updated_at":{
//        "type": "date",
//        "null_value": "null",
//        "format": "[yyyy-MM-dd HH:mm:ss]"
//      }
//    }
//  }
//}
//`
//}
//
//// IndexExists 索引是否存在
//func (a ArticleModel) IndexExists() bool {
//	exists, err := global.ESClient.
//		IndexExists(a.Index()).
//		Do(context.Background())
//	if err != nil {
//		logrus.Error(err.Error())
//		return exists
//	}
//	return exists
//}
//
//// CreateIndex 创建索引
//func (a ArticleModel) CreateIndex() error {
//	if a.IndexExists() {
//		// 有索引
//		a.RemoveIndex()
//	}
//	// 没有索引
//	// 创建索引
//	createIndex, err := global.ESClient.
//		CreateIndex(a.Index()).
//		BodyString(a.Mapping()).
//		Do(context.Background())
//	if err != nil {
//		logrus.Error("创建索引失败")
//		logrus.Error(err.Error())
//		return err
//	}
//	if !createIndex.Acknowledged {
//		logrus.Error("创建失败")
//		return err
//	}
//	logrus.Infof("索引 %s 创建成功", a.Index())
//	return nil
//}
//
//// RemoveIndex 删除索引
//func (a ArticleModel) RemoveIndex() error {
//	logrus.Info("索引存在，删除索引")
//	// 删除索引
//	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
//	if err != nil {
//		logrus.Error("删除索引失败")
//		logrus.Error(err.Error())
//		return err
//	}
//	if !indexDelete.Acknowledged {
//		logrus.Error("删除索引失败")
//		return err
//	}
//	logrus.Info("索引删除成功")
//	return nil
//}
//
//// Create 添加的方法
//func (a *ArticleModel) Create() (err error) {
//	indexResponse, err := global.ESClient.Index().
//		Index(a.Index()).
//		BodyJson(a).Do(context.Background())
//	if err != nil {
//		logrus.Error(err.Error())
//		return err
//	}
//	a.ID = indexResponse.Id
//	return nil
//}
//
//// ISExistData 是否存在该文章
//func (a ArticleModel) ISExistData() bool {
//	res, err := global.ESClient.
//		Search(a.Index()).
//		Query(elastic.NewTermQuery("keyword", a.Title)).
//		Size(1).
//		Do(context.Background())
//	if err != nil {
//		logrus.Error(err.Error())
//		return false
//	}
//	if res.Hits.TotalHits.Value > 0 {
//		return true
//	}
//	return false
//}
//func (a *ArticleModel) GetDataByID(id string) error {
//	res, err := global.ESClient.
//		Get().
//		Index(a.Index()).
//		Id(id).
//		Do(context.Background())
//	if err != nil {
//		return err
//	}
//	err = json.Unmarshal(res.Source, a)
//	return err
//}
