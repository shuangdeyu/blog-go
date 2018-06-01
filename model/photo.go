package model

import (
	"github.com/cihub/seelog"
)

type TpPhoto struct {
	Id      int    `gorm:"type:int(11);AUTO_INCREMENT" json:"id"` // 作品集.摄影 id
	Content string `gorm:"type:text" json:"content"`              // 内容描述
	Type    int    `gorm:"type:tinyint(3);default:0" json:"type"` // 类型，0山，1人，2仙
}

type ShanList struct {
	Content string `json:"content"`
	PicUrl  string `json:"pic_url"`
}

func GetShanList() (list []*ShanList, err error) {
	err = DB.Table("tp_photo p").Select("p.content,i.pic_url").
		Joins("join tp_images i on p.id = i.target_id and i.target = 0").
		Where("p.type = 0").Find(&list).Error
	if err != nil {
		seelog.Error("数据库查询出错", err)
		return nil, err
	}
	return list, err
}

type ImgConList struct {
	Content string
	ImgList []*TpImages
}

func GetImgConList(t int) (list *ImgConList, err error) {
	var tpPhoto []*TpPhoto
	err = DB.Where("type = ?",t).First(&tpPhoto).Error
	if err != nil {
		return nil, err
	}

	var images []*TpImages
	err = DB.Where("target_id = ? and target = 0", tpPhoto[0].Id).Find(&images).Error
	if err != nil {
		return nil, err
	}

	list = new(ImgConList)
	list.Content = tpPhoto[0].Content
	list.ImgList = images
	return list, err
}

func GetImgList() (list []*TpImages, err error) {
	err = DB.Where("target = 1").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, err
}