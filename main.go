package main

import (
	"github.com/daddydemir/MQ-Backend/config"
	"github.com/daddydemir/MQ-Backend/handler"
	"net/http"
)

func main() {
	config.GetConnect()
	println("http://localhost:2121/test")

	server := &http.Server{
		Addr:    ":2121",
		Handler: handler.MainRouting(),
	}

	_ = server.ListenAndServe()
}
