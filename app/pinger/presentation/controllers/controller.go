package controllers

import (
	"encoding/json"
	"net/http"

	"gitlab.studionx.ru/id/pinger/app/pinger/presentation/models"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (s *Controller) HandlePing(w http.ResponseWriter, r *http.Request) {
	response := models.NewPingResponseModel("pong")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
