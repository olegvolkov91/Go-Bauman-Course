package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/models"
	"net/http"
	"strconv"
)

// GetAllArticles ... Возвращает все статьи из БД
func (h *Handler) GetAllArticles(writer http.ResponseWriter, _ *http.Request) {
	initHeaders(writer)
	// Логируем момент начала обработки
	h.api.Logger.Info("Get all articles GET /api/v1/articles")
	articles, err := h.api.Storage.Article().SelectAll()

	if err != nil {
		// Обработка ошибки при подключении к БД
		h.api.Logger.Info("Error while Articles.SelectAll : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}
	respond(writer, http.StatusOK, articles)
}

// GetArticleById ... Возвращаем статью по id из БД
func (h *Handler) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	h.api.Logger.Infof("Get article by id GET /api/v1/article/{%d}", id)

	if err != nil {
		h.api.Logger.Info("Troubles while parsing id param : ", err)
		msg := Message{
			Message:    "Error while accessing id param",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		respond(writer, http.StatusBadRequest, msg)
		return
	}

	h.api.Logger.Infof("Get article by id: %d GET /article/{id}", id)
	article, ok, err := h.api.Storage.Article().FindById(id)

	if err != nil {
		// Обработка ошибки при подключении к БД
		h.api.Logger.Info("Error while Articles.FindById : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}

	if !ok {
		// Обработка ответа, если ошибок нет и статья не найдена
		h.api.Logger.Info("Article not found in Article.FindById")
		msg := Message{
			Message:    "Article not found",
			StatusCode: http.StatusNoContent,
			IsError:    false,
		}
		respond(writer, http.StatusNoContent, msg)
		return
	}
	respond(writer, http.StatusOK, article)
}

// CreateArticle ... Создаём новую статью в БД
func (h *Handler) CreateArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	// Логируем момент начала обработки
	h.api.Logger.Infof("Create article POST /api/v1/article")
	var body models.Article

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		// Обработка получения параметров из body
		h.api.Logger.Info("Invalid json received from client : ", err)
		msg := Message{
			Message:    "Provided json is invalid",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		respond(writer, http.StatusBadRequest, msg)
		return
	}

	article, err := h.api.Storage.Article().Create(&body)
	if err != nil {
		// Обработка ошибки при подключении к БД
		h.api.Logger.Info("Error while creating new article | Article.CreateArticle : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}
	respond(writer, http.StatusCreated, article)
}

// DeleteArticleById ... Удаляем статью из БД
func (h *Handler) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	h.api.Logger.Infof("Delete article by id DELETE /api/v1/article/{%d}", id)

	if err != nil {
		h.api.Logger.Info("Troubles while parsing id param : ", err)
		msg := Message{
			Message:    "Error while accessing id param",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		respond(writer, http.StatusBadRequest, msg)
		return
	}

	h.api.Logger.Infof("Get article by id: %d GET /article/{id}", id)
	article, err := h.api.Storage.Article().DeleteById(id)

	if err != nil {
		// Обработка ошибки при подключении к БД
		h.api.Logger.Info("Error while Articles.DeleteById : ", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}
	// Здесь можно также вернуть просто хедер со статусом 204, но тогда не стоит возвращать никаких данных
	respond(writer, http.StatusOK, article)
}
