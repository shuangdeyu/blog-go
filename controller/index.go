package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"jyj-go/model"
)

// 主页
func HomePage(c *gin.Context){
	c.HTML(http.StatusOK, "index/index.html", gin.H{})
}

// 我的介绍
func CardPage(c *gin.Context) {
	c.HTML(http.StatusOK,"index/card.html", nil)
}

// 微博
func WeiboPage(c *gin.Context) {
	page := c.Param("page")
	if page == "" {
		page = "0"
	}
	// 获取数据
	count := model.WeiboPageSize // 每页总数
	list, err := model.GetWeiboList(page, count)
	// 获取分页数据
	pagination, err := CreatePage(list.Total, count, page, "/weibo")
	//fmt.Println(list.List)
	if err != nil {
		Handle404(c)
	} else {
		c.HTML(http.StatusOK, "index/weibo.html", gin.H{
			"list": list.List,
			"page": pagination,
		})
	}
}

// 404
func Handle404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "base/404.html", nil)
}