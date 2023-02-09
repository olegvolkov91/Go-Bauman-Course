package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port string = "8080"
)

func main() {
	log.Println("Starting pizza rest api ...")
	router := mux.NewRouter()
	router.HandleFunc("/pizza", GetAllPizzas).Methods(http.MethodGet)
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods(http.MethodGet)
	log.Println("configured successfully")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func GetPizzaById(w http.ResponseWriter, req *http.Request) {
}

func GetAllPizzas(w http.ResponseWriter, req *http.Request) {
}
