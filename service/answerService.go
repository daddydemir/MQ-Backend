package service

import (
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/core"
	"github.com/daddydemir/MQ-Backend/model"
	"net/http"
)

func AddAnswer(answer model.Answer) (int, interface{}) {
	result := config.DB.Create(&answer)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	return http.StatusCreated, core.SendMessage(core.Created)
}

func GetAnswersByQuestionId(id string) (int, interface{}) {
	var answers []model.Answer
	result := config.DB.Find(&answers, "question_id = ?", id)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	send := core.SendMessage(core.Ok)
	send["data"] = answers
	return http.StatusOK, send
}
