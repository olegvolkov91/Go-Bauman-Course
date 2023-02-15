package main

import (
	// GORM в 2 - 2.5 раза просаживает скорость работы с БД
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/gin-gorm/models"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/gin-gorm/routers"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/gin-gorm/storage"
	"log"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "host=...user=...password=...dbname=...sslmode=disable")
	if err != nil {
		log.Println("error while accessing database")
	}
	defer storage.DB.Close()
	storage.DB.AutoMigrate(&models.Article{})

	r := routers.SetupRouter()

	// r - gin маршрутизатор
	r.Run()
}
