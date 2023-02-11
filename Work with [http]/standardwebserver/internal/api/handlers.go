package api

import (
	"encoding/json"
	"fmt"
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
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	api.logger.Infof("Get article by id GET /api/v1/article/{%d}", id)

	if err != nil {
		api.logger.Info("Troubles while parsing id param : ", err)
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
	api.logger.Infof("Create article POST /api/v1/article")
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
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	api.logger.Infof("Delete article by id DELETE /api/v1/article/{%d}", id)

	if err != nil {
		api.logger.Info("Troubles while parsing id param : ", err)
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
	initHeaders(writer)
	api.logger.Info("Register user POST /api/v1/register")

	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		api.logger.Info("Invalid json received from client")
		msg := Message{
			Message:    "Provided json is invalid",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	// Находим пользователя в базе
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	// Если такой пользователь есть, то регистрацию не делаем
	if ok {
		api.logger.Info("User with that login already exists")
		msg := Message{
			Message:    "User already exists",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	newUser, err := api.storage.User().Create(&user)

	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		Message:    fmt.Sprintf("User {login: %s}successfully registred!", newUser.Login),
		StatusCode: http.StatusCreated,
		IsError:    false,
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(msg)
}
