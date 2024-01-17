package models

import "gorm.io/gorm"

type PaintingInfo struct {
	gorm.Model
	ImgUrl string `json:"imgUrl" gorm:"size:256"`
}

// AuthorList 作者列表
type AuthorList struct {
	gorm.Model
	AuthorName string `json:"authorName" gorm:"size:32"`
}

// ContentType 内容分类
type ContentType struct {
	gorm.Model
	Name        string `json:"name"`
	ContentType int    `json:"contentType"` // 1.山水 2.花鸟 3.人物
}
