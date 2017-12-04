package handler

import (
	"encoding/json"
	"fmt"
	"go_crud_app/domain"
	"go_crud_app/service"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error:%s", err.Error()), http.StatusBadRequest)
		return
	}
	err = service.InsertUserData(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
