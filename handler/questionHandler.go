package handler

import (
	"encoding/json"
	"github.com/daddydemir/MQ-Backend/model"
	"github.com/daddydemir/MQ-Backend/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func addQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	var question model.Question
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &question)

	code, message := service.AddQuestion(question, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	code, message := service.GetQuestions(t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func getQuestionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := service.GetQuestionById(key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
