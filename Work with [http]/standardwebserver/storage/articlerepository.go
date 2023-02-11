package storage

import (
	"fmt"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/models"
	"log"
)

// ArticleRepository ... Instance of Article Repository (model interface)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	row := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content)

	if err := row.Scan(&a.Id); err != nil {
		return nil, err
	}

	return a, nil
}

func (ar *ArticleRepository) FindById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()

	var founded bool
	if err != nil {
		return nil, founded, err
	}

	var articleFound *models.Article
	for _, a := range articles {
		if a.Id == id {
			articleFound = a
			founded = true
			break
		}
	}

	return articleFound, founded, nil
}

func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, ok, err := ar.FindById(id)

	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)
		if _, err := ar.storage.db.Exec(query, article.Id); err != nil {
			return nil, err
		}
	}
	return article, nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)

	rows, err := ar.storage.db.Query(query)
	if err != nil {
		log.Println("ERROR when SELECT ALL db.Query")
		return nil, err
	}
	defer rows.Close()

	articles := make([]*models.Article, 0)

	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.Id, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}
