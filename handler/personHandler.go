package handler

import (
	"encoding/json"
	"github.com/daddydemir/MQ-Backend/model"
	"github.com/daddydemir/MQ-Backend/service"
	"io/ioutil"
	"net/http"
)

func addPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var person model.Person
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &person)

	code, message := service.PersonAdd(person)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
