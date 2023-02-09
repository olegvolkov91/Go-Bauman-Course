package main

import (
	"log"
	"net/http"
)

var (
	port string = "8080"
)

func main() {
	log.Println("Starting pizza rest api ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
