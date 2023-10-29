package article_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"time"
)

type articleCalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type calendarBuckets struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"calendarBuckets"`
}

// ArticleCalendar es文章按时间聚合
func (ArticleApi) ArticleCalendar(c *gin.Context) {
	// 计算当前时间和上一年度的起始日期
	currentTime := time.Now()
	lastYearStartDate := currentTime.AddDate(-1, 0, 0) // 加1秒以避免包括上一年的最后一刻
	format := "2006-01-02 15:04:05"

	// 创建日期范围查询
	dateRangeQuery := elastic.NewRangeQuery("created_at").
		Gte(lastYearStartDate.Format(format)).
		Lte(currentTime.Format(format))

	// 创建日期直方图聚合
	dateHistogramAgg := elastic.NewDateHistogramAggregation().
		Field("created_at").    // 指定日期字段
		CalendarInterval("day") // 设置聚合间隔为天

	searchResult, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(dateRangeQuery).
		Aggregation("calendar", dateHistogramAgg).
		//Size(0). // 我们只需要聚合结果，不需要文档
		Do(context.Background())
	if err != nil {
		res.Fail(c, res.FAIL_OPER, err.Error())
		return
	}
	var buckets calendarBuckets
	json.Unmarshal(searchResult.Aggregations["calendar"], &buckets)
	searchMap := make(map[string]int)
	for _, item := range buckets.Buckets {
		t, _ := time.Parse(format, item.KeyAsString)
		searchMap[t.Format("2006-01-02")] = item.DocCount
	}

	// 获取今天的日期
	today := time.Now()
	// 计算一年前的今天
	oneYearAgo := today.AddDate(-1, 0, 0)
	var resList []articleCalendarResponse
	for currentDate := oneYearAgo; currentDate.Before(today) || currentDate.Equal(today); currentDate = currentDate.AddDate(0, 0, 1) {
		date := currentDate.Format("2006-01-02")
		response := articleCalendarResponse{
			Date:  date,
			Count: searchMap[date],
		}
		resList = append(resList, response)
	}
	res.OkWithData(c, resList)
}
