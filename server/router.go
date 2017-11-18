package server

import (
	"crud-app/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/user", handler.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", handler.GetUserHandler).Methods("GET")
	return router
}
