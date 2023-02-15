package helpers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

func RespondJSON(ctx *gin.Context, statusCode int, data interface{}) {
	log.Println("status code: ", statusCode)

	var message Message
	message.StatusCode = statusCode
	message.Data = data
	ctx.JSON(statusCode, message)
}
