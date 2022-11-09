package service

import (
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/core"
	"github.com/daddydemir/MQ-Backend/model"
	"github.com/daddydemir/MQ-Backend/security"
	"net/http"
)

func PersonAdd(person model.Person) (int, interface{}) {
	if person.Nickname == "" {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	var temp model.Person

	r := config.DB.Find(&temp, "email = ?", person.Email)
	if r.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	person.Password = security.HashPassword(person.Password)
	result := config.DB.Create(&person)

	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Created)
		send["data"] = person
		return http.StatusCreated, send
	}
}
