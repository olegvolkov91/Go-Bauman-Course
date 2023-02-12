package handler

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

// respond ... Вспомогательная функция для возврата ответа пользователю
func respond(writer http.ResponseWriter, code int, data interface{}) {
	writer.WriteHeader(code)
	if data != nil {
		json.NewEncoder(writer).Encode(data)
	}
}
