package api

import (
	"encoding/json"
	"net/http"
)

// Message ... Вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// GetAllArticles ... Возвращает все статьи из БД
func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	// Логируем момент начала обработки
	api.logger.Info("Get all articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()

	if err != nil {
		// Обработка ошибки при подключении к БД
		api.logger.Info("Error while Articles.SelectAll : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {

}

func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {

}

func (api *API) CreateArticle(writer http.ResponseWriter, req *http.Request) {

}

func (api *API) RegisterUser(writer http.ResponseWriter, req *http.Request) {

}
