package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodos(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/alltodos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []todo
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "1", response[0].ID)
	assert.Equal(t, "Wash dishes", response[0].Title)
	assert.Equal(t, false, response[0].Completed)
	assert.Equal(t, "2", response[1].ID)
	assert.Equal(t, "Do laundry", response[1].Title)
	assert.Equal(t, false, response[1].Completed)
	assert.Equal(t, "3", response[2].ID)
	assert.Equal(t, "Clean room", response[2].Title)
	assert.Equal(t, false, response[2].Completed)
}
