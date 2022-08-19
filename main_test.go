package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/pkg/data"

	"github.com/go-playground/assert/v2"
)

func TestGetRequest(t *testing.T) {
	router := setupRouter()
	responsdata := data.Todo
	httpwriter := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/todos", nil)

	router.ServeHTTP(httpwriter, request)

	var jsonBody []data.Tododata

	err := json.Unmarshal(httpwriter.Body.Bytes(), &jsonBody)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, httpwriter.Code)
	assert.Equal(t, responsdata, jsonBody)
}
