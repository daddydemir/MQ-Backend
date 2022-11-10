package handler

import (
	"encoding/json"
	"github.com/daddydemir/MQ-Backend/model"
	"github.com/daddydemir/MQ-Backend/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func addAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	var answer model.Answer
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &answer)

	code, message := service.AddAnswer(answer, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func getAnswersByQuestionId(w http.ResponseWriter, r *http.Request) {
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

	code, message := service.GetAnswersByQuestionId(key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
