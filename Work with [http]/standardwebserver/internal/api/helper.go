package api

import (
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
	api.logger.SetLevel(logLevel)
	return nil
}

func (api *API) configureRouter() {
	api.router.HandleFunc(prefix+"/articles", api.GetAllArticles).Methods("GET")
	api.router.HandleFunc(prefix+"/articles/{id}", api.GetArticleById).Methods(http.MethodGet)
	api.router.HandleFunc(prefix+"/articles/{id}", api.DeleteArticleById).Methods(http.MethodDelete)
	api.router.HandleFunc(prefix+"/articles", api.CreateArticle).Methods(http.MethodPost)

	api.router.HandleFunc(prefix+"/user/register", api.RegisterUser).Methods(http.MethodPost)
}

func (api *API) configureStorage() error {
	// Создаём экземпляр хранилища
	store := storage.New(api.config.Storage)
	// Пробуем установить соединение с хранилищем, если невозможно - возвращаем ошибку
	if err := store.Open(); err != nil {
		return err
	}
	// сохраняем его в экземпляре API для дальнейшей работы
	api.storage = store
	return nil
}
