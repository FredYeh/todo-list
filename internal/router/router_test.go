package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FredYeh/todo-list/internal/store/items"
	"github.com/stretchr/testify/assert"
)

func TestRouters(t *testing.T) {
	router := Router("test")
	// Testing POST
	task := items.Task{Name: "Testing", Status: 0}
	reqBody, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Testing GET
	req, _ = http.NewRequest("GET", "/tasks", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	body := make([]map[string]any, 0)
	err := json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Testing", body[0]["name"])
	assert.Equal(t, "0", body[0]["status"])
	assert.Len(t, body, 1)

	// Testing PUT
	task = items.Task{Name: "Testing", Status: 1}
	reqBody, err = json.Marshal(task)
	assert.Equal(t, err, nil)
	req, _ = http.NewRequest("PUT", "/tasks/"+body[0]["id"].(string), bytes.NewBuffer(reqBody))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Testing GET
	req, _ = http.NewRequest("GET", "/tasks", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	body = make([]map[string]any, 0)
	err = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Testing", body[0]["name"])
	assert.Equal(t, "1", body[0]["status"])
	assert.Len(t, body, 1)

	// Testing DELETE
	req, _ = http.NewRequest("DELETE", "/tasks/"+body[0]["id"].(string), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Testing GET
	req, _ = http.NewRequest("GET", "/tasks", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	body = make([]map[string]any, 0)
	err = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Len(t, body, 0)
}
