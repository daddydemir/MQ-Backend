package handler

import (
	"encoding/json"
	"github.com/daddydemir/MQ-Backend/auth"
	"github.com/daddydemir/MQ-Backend/model"
	"github.com/daddydemir/MQ-Backend/service"
	"io/ioutil"
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var content map[string]string
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &content)

	code, message := auth.Login(content["Mail"], content["Password"])
	if code == 200 {
		var history model.History
		var person model.Person
		person = message["data"].(model.Person)
		history.Time = time.Now()
		history.IpAddress = r.RemoteAddr
		history.PersonId = person.Id

		c, m := service.AddHistory(history)
		if c != 201 {
			w.WriteHeader(c)
			_ = json.NewEncoder(w).Encode(m)
		}
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
