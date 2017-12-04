package handler

import (
	"encoding/json"
	"go_crud_app/service"
	"net/http"
)

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {

	user, err := service.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
