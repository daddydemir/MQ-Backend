package handler

import (
	"github.com/daddydemir/MQ-Backend/core"
	"net/http"
)

func _tokenCheck(r *http.Request) (bool, interface{}, string) {
	token := r.Header["Authorization"]
	if token == nil {
		return false, core.SendMessage(core.NotLogin), ""
	} else {
		return true, nil, token[0]
	}
}
