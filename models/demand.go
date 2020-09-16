package models

import "gorm.io/gorm"

type Demand struct {
	gorm.Model
	// 标题
	Title string
	// 内容
	Body string `gorm:"type:text"`
	// 照片
	Photo string `gorm:"size:1024"`
	// 价格,单位是分
	Price uint
	// 类型(0: 代吃 1:代做)
	Style uint
	// 省
	Province string
	// 市
	City string
	//区
	Area string
	// 地址
	Address string `gorm:"size:1024"`
	// 热度
	Heat uint `gorm:"default:0"`
	// 创建人用户id
	CreateUserId uint  `gorm:"default:null"`
	Users        Users `gorm:"foreignKey:CreateUserId"`
}

func AddDemand(demand Demand) bool {
	DB.Create(&demand)
	return true
}
