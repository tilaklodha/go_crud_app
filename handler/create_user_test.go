package handler

import (
	"bytes"
	"go_crud_app/appcontext"
	"go_crud_app/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTest() {
	config.Load()
	appcontext.Initiate()
}

func TestCreateUserWithInvalidBody(t *testing.T) {
	setupTest()
	rr := httptest.NewRecorder()
	var requestBody = []byte(`"{"first_name":"abc", "last_name":"xyz", "city":"city"}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(requestBody))

	CreateUserHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateUserWithValidBody(t *testing.T) {
	setupTest()
	rr := httptest.NewRecorder()
	var requestBody = []byte(`{"first_name":"abc", "last_name":"xyz", "city":"city"}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(requestBody))

	CreateUserHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func cleanUpDB() {
	db := appcontext.GetDB()
	db.MustExec("DELETE from users")
}
