package service

import (
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/core"
	"github.com/daddydemir/MQ-Backend/model"
	"net/http"
)

func AddQuestion(question model.Question) (int, interface{}) {
	result := config.DB.Create(&question)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	return http.StatusCreated, core.SendMessage(core.Created)
}

func GetQuestions() (int, interface{}) {
	var questions []model.Question
	result := config.DB.Find(&questions)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	send := core.SendMessage(core.Ok)
	send["data"] = questions
	return http.StatusOK, send
}

func GetQuestionById(id string) (int, interface{}) {
	var question model.Question
	result := config.DB.Find(&question, "id = ?", id)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	send := core.SendMessage(core.Ok)
	send["data"] = question
	return http.StatusOK, send
}
