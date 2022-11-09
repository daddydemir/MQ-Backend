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

	handler := cors.AllowAll().Handler(r)
	return handler
}
func test(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("golang")
}
