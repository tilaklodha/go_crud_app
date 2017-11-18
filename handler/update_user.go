package handler

import (
	"crud-app/domain"
	"crud-app/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["id"])
	var user domain.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusBadRequest)
	}

	if userNotPresent(userId) {
		http.Error(w, fmt.Sprintf("User does not exist"), http.StatusBadRequest)
		return
	}

	updateUser, err := service.UpdateUser(&user, userId)

	response, err := json.Marshal(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func userNotPresent(userId int) bool {
	user, err := service.GetUser(userId)
	if err != nil {
		fmt.Errorf("Error checking user in database")
	}

	if *user == (domain.User{}) {
		return true
	}
	return false
}
