package article_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"blog_gin/utils/jwts"
	"context"
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"strings"
	"time"
	"unicode/utf8"
)

type articleUpdateRequest struct {
	ID       string `json:"id,omitempty" validate:"required" label:"文章ID"`
	Title    string `json:"title"  label:"文章标题"`             // 文章标题
	Abstract string `json:"abstract"`                            // 文章简介
	Content  string `json:"content,omitempty"  label:"文章内容"` // 文章内容

	Category string `json:"category,omitempty"` // 文章分类
	Source   string `json:"source,omitempty"`   // 文章来源
	Link     string `json:"link,omitempty"`     // 原文链接

	BannerID uint        `json:"banner_id,omitempty"` // 文章封面id
	Tags     ctype.Array `json:"tags,omitempty"`      // 文章标签
}

type EsUpdateArticleModel struct {
	UpdatedAt string `json:"updated_at,omitempty"` // 更新时间

	Title    string `json:"title,omitempty"`    // 文章标题
	Keyword  string `json:"keyword,omitempty"`  // 文章标题
	Abstract string `json:"abstract,omitempty"` // 文章简介
	Content  string `json:"content,omitempty"`  // 文章内容

	LookCount     int `json:"look_count,omitempty"`     // 浏览量
	CommentCount  int `json:"comment_count,omitempty"`  // 评论量
	DiggCount     int `json:"digg_count,omitempty"`     // 点赞量
	CollectsCount int `json:"collects_count,omitempty"` // 收藏量

	UserNickName string `json:"user_nick_name,omitempty"` //用户昵称
	UserAvatar   string `json:"user_avatar,omitempty"`    // 用户头像

	Category string `json:"category,omitempty"`          // 文章分类
	Source   string `json:"source,omit(list),omitempty"` // 文章来源
	Link     string `json:"link,omit(list),omitempty"`   // 原文链接

	BannerID  uint   `json:"banner_id,omitempty"`  // 文章封面id
	BannerUrl string `json:"banner_url,omitempty"` // 文章封面

	Tags ctype.Array `json:"tags,omitempty"` // 文章标签
}

// ArticleList 文章列表
func (ArticleApi) ArticleUpdate(c *gin.Context) {
	var cr articleUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}
	model := models.ArticleModel{}
	// 获取esUserID

	searchResult, err := global.ESClient.Get().
		Index(model.Index()).
		Id(cr.ID). // 查询条件
		Do(context.Background())
	if err != nil {
		res.Fail(c, 10004, err.Error())
		return
	}

	json.Unmarshal(searchResult.Source, &model)
	// 判断当前token的userID 是否与该文章的userID 相同
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	userID := payload.UserID
	if userID != model.UserID {
		res.Fail(c, 4002, "无权限")
		return
	}

	userNickName := payload.NickName

	// 将markdown转为html
	if cr.Content != "" {
		unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
		// xss过滤
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
		if nodes := doc.Find("script").Nodes; len(nodes) > 0 {
			doc.Find("script").Remove()
			converter := md.NewConverter("", true, nil)
			html, _ := doc.Html()
			markdown, _ := converter.ConvertString(html)
			cr.Content = markdown
		}
		// 如果简介为空, 简介就提取正文前30个字符
		if cr.Abstract == "" {
			// 计算字符串长度
			length := utf8.RuneCountInString(doc.Text())
			if length >= 100 {
				// 如果长度大于100，则截取前100个字符
				cr.Abstract = doc.Text()[:100]
			} else {
				cr.Abstract = doc.Text()
			}
		}
	}
	// 根据banner_id 获得banner.path
	var bannerUrl string
	if cr.BannerID != 0 {
		result := global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl)
		err = result.Error
		if err != nil {
			res.Fail(c, res.FAIL_OPER, err.Error())
			return
		}
		if result.RowsAffected == 0 {
			res.Fail(c, res.FAIL_OPER, "todo banner不存在")
			return
		}
	}

	// 用户头像
	var userAvatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&userAvatar).Error
	if err != nil {
		res.Fail(c, res.FAIL_OPER, "todo user不存在")
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	article := EsUpdateArticleModel{
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserNickName: userNickName,
		UserAvatar:   userAvatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	//EsUpdateArticleMap := structs.Map(&article)
	fmt.Println(article)

	//global.ESClient.Update().
	//	Index(model.Index()).
	//	Id(cr.ID).
	//	Refresh("true").
	//	Doc(EsUpdateArticleMap).
	//	Do(context.Background())
	//if err != nil {
	//	res.Fail(c, 10001, err.Error())
	//	return
	//}
	res.OkWithData(c, article)
}
