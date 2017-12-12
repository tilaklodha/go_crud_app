package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		log.Fatalf("Failed to create new request")
		return
	}
	PingHandler(responseRecorder, request)
	response := responseRecorder.Result()
	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("ping handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	want := `"ping":"pong"`
	body, _ := ioutil.ReadAll(response.Body)
	got := string(body)
	if got != want {
		t.Errorf("ping handler returned unexpected body: got %v want %v", got, want)
	}

}
