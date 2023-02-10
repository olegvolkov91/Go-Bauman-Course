package utils

import (
	"github.com/gorilla/mux"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/books_api/handlers"
	"net/http"
)

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc(prefix, handlers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods(http.MethodPut)
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookById).Methods(http.MethodDelete)
}

func BuildManyBooksResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods(http.MethodGet)
}
