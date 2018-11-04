package test

// HTTP testing with Go

import (
	"awesomeProject/functions"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router(path string, crudMethod string, f func(http.ResponseWriter, *http.Request)) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(path, functions.GetPeople).Methods(crudMethod)
	return router
}

func TestGetPeople(t *testing.T) {
	request, _ := http.NewRequest("GET", "/people", nil)
	response := httptest.NewRecorder()
	Router("/people", "GET", functions.GetPeople).ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
