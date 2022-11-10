package handler

import (
	"encoding/json"
	"github.com/daddydemir/MQ-Backend/service"
	"github.com/gorilla/mux"
	"net/http"
)

func getByPersonId(w http.ResponseWriter, r *http.Request) {
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

	code, message := service.GetByPersonId(key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
