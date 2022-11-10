package handler

import (
	"encoding/json"
	"github.com/daddydemir/MQ-Backend/auth"
	"io/ioutil"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var content map[string]string
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &content)

	code, message := auth.Login(content["Mail"], content["Password"])
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
