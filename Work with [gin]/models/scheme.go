package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	//ID uint `gorm:"primary_key" json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

// TableName ... Требование GORM - наличие метода возвращающего имя таблицы
func (a *Article) TableName() string {
	return "article"
}
