package api

import (
	"github.com/gorilla/mux"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/handler"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

// API ... Base api service description
type API struct {
	config *Config
	Logger *logrus.Logger
	router *mux.Router
	// Добавление поля для работы с хранилищем
	Storage *storage.Storage
	// Добавление обработчиков
	handler *handler.Handler
}

// New ... Build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		Logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ... HTTP server/configure loggers
func (api *API) Start() error {
	// Конфигурируем логгер
	if err := api.configureLogger(); err != nil {
		return err
	}
	api.Logger.Infof("starting api server at port %s", api.config.Port)

	// Конфигурируем хендлеры
	api.configureHandler()
	// Конфигурируем роутер
	api.configureRouter()
	// Конфигурируем хранилище
	if err := api.configureStorage(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.Port, api.router)
}
