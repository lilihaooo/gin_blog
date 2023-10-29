package article_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"blog_gin/utils"
	"blog_gin/utils/jwts"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

type articleRequest struct {
	Title    string `json:"title" validate:"required" label:"文章标题"`   // 文章标题
	Abstract string `json:"abstract"`                                 // 文章简介
	Content  string `json:"content" validate:"required" label:"文章内容"` // 文章内容

	Category string `json:"category"` // 文章分类
	Source   string `json:"source"`   // 文章来源
	Link     string `json:"link"`     // 原文链接

	BannerID uint `json:"banner_id"` // 文章封面id

	Tags ctype.Array `json:"tags"` // 文章标签
}

func (ArticleApi) ArticleCreate(c *gin.Context) {
	var cr articleRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.Fail(c, res.INVALID_PARAMS, "")
		return
	}
	if vRes := utils.ZhValidate(&cr); vRes != nil {
		res.FailValidate(c, vRes)
		return
	}

	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	userID := payload.UserID
	userNickName := payload.NickName
	/*
			校验content xss攻击
		1. 先将markdown转为html
			github.com/russross/blackfriday
		2. 检测里面有没有<scrip>标签, 如果有将其删除 (html获取文本内容，xss过滤)
			github.com/PuerkitoBio/goquery
		3. 再转为markdown文本
			github.com/JohannesKaufmann/html-to-markdown
	*/
	// 将markdown转为html
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
	// 如果banner_id 为0 , 就随机获取一个banner_id
	if cr.BannerID == 0 {
		var bannerIDList []uint
		if err := global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList).Error; err != nil {
			res.Fail(c, res.FAIL_OPER, err.Error())
			return
		}
		if len(bannerIDList) == 0 {
			res.Fail(c, res.FAIL_OPER, "todo 没有banner数据")
			return
		}
		// 从[]uint中随机获取一个
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(bannerIDList)) // 初始化随机数生成器
		randomElement := bannerIDList[randomIndex]  // 获取随机元素
		cr.BannerID = randomElement
	}

	// 根据banner_id 获得banner.path
	var bannerUrl string
	result := global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl)
	err := result.Error
	if err != nil {
		res.Fail(c, res.FAIL_OPER, err.Error())
		return
	}
	if result.RowsAffected == 0 {
		res.Fail(c, res.FAIL_OPER, "todo banner不存在")
		return
	}

	// 用户头像
	var userAvatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&userAvatar).Error
	if err != nil {
		res.Fail(c, res.FAIL_OPER, "todo user不存在")
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   userAvatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	// 判断文章标题是否存在
	exist := article.ISExistData()
	if exist {
		res.Fail(c, res.FAIL_OPER, "文章已经存在")
		return
	}

	err = article.Create()
	if err != nil {
		res.Fail(c, res.FAIL_OPER, "todo 文章发布失败")
	}
	res.Ok(c)

}
