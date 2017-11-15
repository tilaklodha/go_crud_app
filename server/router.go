package server

import (
	"crud-app/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	return router
}
