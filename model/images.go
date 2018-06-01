package model

type TpImages struct {
	Id       int    `gorm:"type:int(11);AUTO_INCREMENT" json:"id"`   // id
	PicUrl   string `gorm:"type:varchar(255)" json:"pic_url"`        // 图片地址
	TargetId int    `gorm:"type:int(11);default:0" json:"target_id"` // 目标id
	Target   int    `gorm:"type:tinyint(3);default:0" json:"target"` // 目标类型，0摄影，1相册
}
