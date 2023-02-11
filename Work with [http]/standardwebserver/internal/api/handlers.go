package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/models"
	"net/http"
	"strconv"
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
func (api *API) GetAllArticles(writer http.ResponseWriter, _ *http.Request) {
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

// GetArticleById ... Возвращаем статью по id из БД
func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		api.logger.Info("GetArticleById : ", err)
		msg := Message{
			Message:    "Error while accessing id param",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	api.logger.Infof("Get article by id: %d GET /article/{id}", id)
	article, ok, err := api.storage.Article().FindById(id)

	if err != nil {
		// Обработка ошибки при подключении к БД
		api.logger.Info("Error while Articles.FindById : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		// Обработка ответа, если ошибок нет и статья не найдена
		api.logger.Info("Article not found in Article.FindById")
		msg := Message{
			Message:    "Article not found",
			StatusCode: http.StatusNoContent,
			IsError:    false,
		}
		writer.WriteHeader(http.StatusNoContent)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)
}

// CreateArticle ... Создаём новую статью в БД
func (api *API) CreateArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	// Логируем момент начала обработки
	api.logger.Info("Post article POST /api/v1/articles")
	var body models.Article

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		// Обработка получения параметров из body
		api.logger.Info("Invalid json received from client : ", err)
		msg := Message{
			Message:    "Provided json is invalid",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	article, err := api.storage.Article().Create(&body)
	if err != nil {
		// Обработка ошибки при подключении к БД
		api.logger.Info("Error while creating new article | Article.CreateArticle : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(article)
}

// DeleteArticleById ... Удаляем статью из БД
func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		api.logger.Info("DeleteArticleById : ", err)
		msg := Message{
			Message:    "Error while accessing id param",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	api.logger.Infof("Get article by id: %d GET /article/{id}", id)
	article, err := api.storage.Article().DeleteById(id)

	if err != nil {
		// Обработка ошибки при подключении к БД
		api.logger.Info("Error while Articles.DeleteById : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Здесь можно также вернуть просто хедер, но тогда не стоит возвращать никаких данных
	// writer.WriteHeader(http.StatusNoContent)

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)
}

func (api *API) RegisterUser(writer http.ResponseWriter, req *http.Request) {

}
