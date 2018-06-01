package model

import (
	"strconv"
	"github.com/cihub/seelog"
)

type TpWeibo struct {
	Id      int    `gorm:"type:int(6);AUTO_INCREMENT" json:"id"`  // 微博id
	Name    string `gorm:"type:varchar(50);not null" json:"name"` // 短语
	Content string `gorm:"type:text;not null" json:"content"`     // 长语
	Link    string `gorm:"type:varchar(255)" json:"link"`         // 外链
}

type WeiboList struct {
	Total   int        `json:"total"`    // 微博总数
	PerPage int        `json:"per_page"` // 下一页页码
	List    []*TpWeibo `json:"list"`     // 微博列表
}

/**
 * 获取微博列表
 */
func GetWeiboList(page string, count int) (list *WeiboList, err error) {
	p, err := strconv.ParseInt(page, 10, 0)
	if err != nil {
		seelog.Error("参数类型错误", err)
		return nil, err
	}
	var weiBo []*TpWeibo
	// 获取微博总数
	var total int
	err = DB.Find(&weiBo).Count(&total).Error
	if err != nil {
		seelog.Error("数据库查询出错", err)
		return nil, err
	}
	// 获取列表
	offset := int(p)*count
	err = DB.Order("id desc").Offset(offset).Limit(count).Find(&weiBo).Error

	list = new(WeiboList)
	list.Total = total
	list.PerPage = int(p) + 1
	list.List = weiBo
	return list, err
}