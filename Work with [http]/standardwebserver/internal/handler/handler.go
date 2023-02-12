package handler

import (
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/api"
)

type Handler struct {
	api *api.API
}

func New(api *api.API) *Handler {
	return &Handler{api}
}
