package models

import (
	_ "github.com/lib/pq"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/gin-gorm/storage"
)

func GetAllArticles(a []Article) error {
	return storage.DB.Find(a).Error
}

func CreateNewArticle(a *Article) error {
	return storage.DB.Create(a).Error
}

func GetArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).First(a).Error
}

func UpdateArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).Update(a).Error
}

func DeleteArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).Delete(a).Error
}
