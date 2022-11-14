package service

import (
	"github.com/daddydemir/MQ-Backend/auth"
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/core"
	"github.com/daddydemir/MQ-Backend/model"
	"net/http"
)

func AddHistory(history model.History) (int, interface{}) {
	result := config.DB.Create(&history)

	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		return http.StatusCreated, core.SendMessage(core.Created)
	}
}

func GetByPersonId(id string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}

	var history []model.History
	result := config.DB.Preload("Person").Find(&history, "person_id = ?", id)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = history
		return http.StatusOK, send
	}
}
