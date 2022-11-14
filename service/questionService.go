package service

import (
	"github.com/daddydemir/MQ-Backend/auth"
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/core"
	"github.com/daddydemir/MQ-Backend/model"
	"net/http"
)

func AddQuestion(question model.Question, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	result := config.DB.Create(&question)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	return http.StatusCreated, core.SendMessage(core.Created)
}

func GetQuestions(token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	var questions []model.Question
	result := config.DB.Find(&questions)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	send := core.SendMessage(core.Ok)
	send["data"] = questions
	return http.StatusOK, send
}

func GetQuestionById(id string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}

	var question model.Question
	result := config.DB.Preload("Person").Find(&question, "id = ?", id)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	send := core.SendMessage(core.Ok)
	send["data"] = question
	return http.StatusOK, send
}
