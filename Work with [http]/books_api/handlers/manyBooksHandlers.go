package handlers

import (
	"encoding/json"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/books_api/models"
	"log"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get info about all books in database")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
