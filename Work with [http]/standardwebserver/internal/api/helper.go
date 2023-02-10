package api

import (
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/storage"
	"github.com/sirupsen/logrus"
	"net/http"
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
	api.router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello"))
	})
}

func (api *API) configureStorage() error {
	// Создаём экземпляр хранилища
	store := storage.New(api.config.Storage)
	// Пробуем установить соединение с хранилищем, если невозможно - возвращаем ошибку
	if err := store.Open(); err != nil {
		return err
	}
	api.storage = store
	return nil
}
