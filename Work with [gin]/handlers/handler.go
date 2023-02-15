package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/gin-gorm/helpers"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/gin-gorm/models"
	"net/http"
)

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	err := models.GetAllArticles(articles)
	if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, articles)
		return
	}
	helpers.RespondJSON(c, http.StatusOK, articles)
}
func GetArticleById(c *gin.Context) {
	id := c.Params.ByName("id")
	var article models.Article

	if err := models.GetArticleById(&article, id); err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, article)
		return
	}
	helpers.RespondJSON(c, http.StatusOK, article)
}

func CreateNewArticle(c *gin.Context) {
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, article)
		return
	}

	if err := models.CreateNewArticle(&article); err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, nil)
		return
	}
	helpers.RespondJSON(c, http.StatusCreated, article)
}

func UpdateArticleById(c *gin.Context) {
	id := c.Params.ByName("id")
	var article models.Article

	if err := models.GetArticleById(&article, id); err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, article)
		return
	}
	c.BindJSON(&article)

	if err := models.UpdateArticleById(&article, id); err != nil {
		helpers.RespondJSON(c, http.StatusInternalServerError, nil)
		return
	}
	helpers.RespondJSON(c, http.StatusAccepted, article)
}

func DeleteArticleById(c *gin.Context) {
	id := c.Params.ByName("id")
	var article models.Article

	if err := models.DeleteArticleById(&article, id); err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, nil)
		return
	}
	helpers.RespondJSON(c, http.StatusAccepted, article)
}
