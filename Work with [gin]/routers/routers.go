package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/gin-gorm/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // аналог mux.NewRouter()
	// В Gin принято группировать ресурсы

	// Указываем префикс
	api := router.Group("/api/v1")
	{
		article := api.Group("/article")
		{
			article.GET("/", handlers.GetAllBooks)
			article.GET("/:id", handlers.GetBookById)
			article.POST("/", handlers.CreateNewBook)
			article.PUT("/:id", handlers.UpdateArticleById)
			article.DELETE("/:id", handlers.DeleteArticleById)
		}
	}
	return router
}
