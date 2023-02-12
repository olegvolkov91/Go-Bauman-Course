package api

import (
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/handler"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/middleware"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	prefix string = "/api/v1"
)

func (api *API) configureLogger() error {
	logLevel, err := logrus.ParseLevel(api.config.LogLevel)
	if err != nil {
		return err
	}
	api.Logger.SetLevel(logLevel)
	return nil
}

func (api *API) configureHandler() {
	api.handler = handler.New(api)
}

func (api *API) configureRouter() {
	api.router.HandleFunc(prefix+"/articles", api.handler.GetAllArticles).Methods(http.MethodGet)
	//api.router.HandleFunc(prefix+"/articles/{id}", api.handler.GetArticleById).Methods(http.MethodGet)
	api.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(http.HandlerFunc(api.handler.GetArticleById)))

	api.router.HandleFunc(prefix+"/articles/{id}", api.handler.DeleteArticleById).Methods(http.MethodDelete)
	api.router.HandleFunc(prefix+"/articles", api.handler.CreateArticle).Methods(http.MethodPost)

	api.router.HandleFunc(prefix+"/user/register", api.handler.RegisterUser).Methods(http.MethodPost)

	api.router.HandleFunc(prefix+"user/auth", api.handler.AuthenticateUser).Methods(http.MethodPost)
}

func (api *API) configureStorage() error {
	// Создаём экземпляр хранилища
	store := storage.New(api.config.Storage)
	// Пробуем установить соединение с хранилищем, если невозможно - возвращаем ошибку
	if err := store.Open(); err != nil {
		return err
	}
	// сохраняем его в экземпляре API для дальнейшей работы
	api.Storage = store
	return nil
}
