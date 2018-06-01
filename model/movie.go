package model

import (
	"time"
	"strconv"
	"github.com/cihub/seelog"
)

type TpMovie struct {
	Id         int       `gorm:"type:mediumint(7);AUTO_INCREMENT" json:"id"`   // 文章id
	Title      string    `gorm:"type:varchar(50);not null" json:"title"`       // 标题
	Cover      string    `gorm:"type:varchar(255);not null" json:"cover"`      // 封面图
	Content    string    `gorm:"type:text;not null" json:"content"`            // 内容
	Time       time.Time `gorm:"type:date;not null" json:"time"`               // 创建日期
	IsMarkdown string    `gorm:"type:varchar(10);not null" json:"is_markdown"` // 是否markdown格式
}

type ArticleList struct {
	Total   int        `json:"total"`    // 文章总数
	PerPage int        `json:"per_page"` // 下一页页码
	List    []*TpMovie `json:"list"`     // 文章列表
}

type ArticleDetail struct {
	Title      string
	Content    string
	Time       time.Time
	IsMarkdown string
}

func GetArticleList(page string, count int) (list *ArticleList, err error) {
	p, err := strconv.ParseInt(page, 10, 0)
	if err != nil {
		seelog.Error("参数类型错误", err)
		return nil, err
	}
	var article []*TpMovie
	// 获取文章总数
	var total int
	err = DB.Find(&article).Count(&total).Error
	if err != nil {
		seelog.Error("数据库查询出错", err)
		return nil, err
	}
	// 获取列表
	offset := int(p)*count
	err = DB.Order("id desc").Offset(offset).Limit(count).Find(&article).Error

	list = new(ArticleList)
	list.Total = total
	list.PerPage = int(p) + 1
	list.List = article
	return list, err
}

func GetArticleDetail(id string) (data *ArticleDetail, err error) {
	var article []*TpMovie
	err = DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		seelog.Error("文章详情查询出错", err)
		return nil, err
	}

	data = new(ArticleDetail)
	data.Title = article[0].Title
	data.Content = article[0].Content
	data.Time = article[0].Time
	data.IsMarkdown = article[0].IsMarkdown
	return data, err
}
