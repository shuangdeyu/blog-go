package controller

import (
	"github.com/gin-gonic/gin"
	"jyj-go/model"
	"net/http"
)

/**
 * 文章列表
 */
func ArticlePage(c *gin.Context) {
	page := c.Param("page")
	if page == "" {
		page = "0"
	}
	// 获取数据
	count := model.ArticlePageSize // 每页总数
	list, err := model.GetArticleList(page, count)
	// 获取分页数据
	pagination, err := CreatePage(list.Total, count, page, "/article")

	if err != nil {
		Handle404(c)
	} else {
		c.HTML(http.StatusOK, "article/index.html", gin.H{
			"list": list.List,
			"page": pagination,
		})
	}
}

/**
 * 文章详情
 */
func DetailPage(c *gin.Context) {
	i := c.Query("i")
	data, err := model.GetArticleDetail(i)
	if err != nil {
		Handle404(c)
	} else {
		c.HTML(http.StatusOK, "article/detail.html", gin.H{
			"title": data.Title,
			"content": data.Content,
			"time": data.Time,
			"is_markdown": data.IsMarkdown,
		})
	}
}
