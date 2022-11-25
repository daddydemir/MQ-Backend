package service

import (
	"github.com/daddydemir/MQ-Backend/auth"
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/core"
	"github.com/daddydemir/MQ-Backend/dao"
	"github.com/daddydemir/MQ-Backend/model"
	"net/http"
	"time"
)

func AddAnswer(answer dao.AnswerAdd, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	var temp model.Answer
	temp.QuestionId = answer.QuestionId
	temp.Content = answer.Content
	temp.UpdateDate = time.Now()
	result := config.DB.Create(&temp)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	return http.StatusCreated, core.SendMessage(core.Created)
}

func GetAnswersByQuestionId(id string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}

	var answers []model.Answer
	result := config.DB.Preload("Question").Find(&answers, "question_id = ?", id)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	send := core.SendMessage(core.Ok)
	send["data"] = answers
	return http.StatusOK, send
}
