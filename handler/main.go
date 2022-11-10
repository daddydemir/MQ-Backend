package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func MainRouting() http.Handler {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/test", test)

	// Person
	r.HandleFunc("/person/add", addPerson).Methods(POST)

	// History
	r.HandleFunc("/history/{id:[0-9]+}", getByPersonId).Methods(GET)

	// Question
	r.HandleFunc("/question/add", addQuestion).Methods(POST)
	r.HandleFunc("/questions", getQuestions).Methods(GET)
	r.HandleFunc("/question/{id:[0-9]+}", getQuestionById).Methods(GET)

	// Answer
	r.HandleFunc("/answer/add", addAnswer).Methods(POST)
	r.HandleFunc("/answer/{id:[0-9]+}", getAnswersByQuestionId).Methods(GET)

	// Auth
	r.HandleFunc("/login", login).Methods(POST)

	handler := cors.AllowAll().Handler(r)
	return handler
}
func test(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(r.RemoteAddr)
}
