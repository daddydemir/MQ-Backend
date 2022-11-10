package auth

import (
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/core"
	"github.com/daddydemir/MQ-Backend/model"
	"github.com/daddydemir/MQ-Backend/security"
	"net/http"
)

func Login(email string, password string) (int, interface{}) {
	var user model.Person
	config.DB.Find(&user, "email = ?", email)
	if user.Id == 0 {
		return http.StatusBadRequest, core.SendMessage(core.LoginFail)
	} else {
		isTrue := security.CheckPassword(user.Password, password)
		if isTrue {
			token := GenerateToken(user.Email)
			send := core.SendMessage(core.Ok)
			send["data"] = user
			send["token"] = token
			return http.StatusOK, send
		}
	}
	return http.StatusBadRequest, core.SendMessage(core.LoginFail)
}
