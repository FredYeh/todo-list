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
	task := items.Task{Name: "Testing", Status: items.Todo}
	reqBody, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	body := make(map[string]any)
	err := json.Unmarshal(w.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "New task created", body["msg"].(string))

	// Testing GET
	req, _ = http.NewRequest("GET", "/tasks", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	body = make(map[string]any)
	err = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	taskMap := body["result"].([]any)[0].(map[string]any)
	assert.Equal(t, "Testing", taskMap["name"].(string))
	assert.Equal(t, "Todo", taskMap["status"].(string))
	assert.Len(t, body, 1)

	// Testing PUT
	task = items.Task{Name: "Testing", Status: items.Done}
	reqBody, err = json.Marshal(task)
	assert.Equal(t, err, nil)
	req, _ = http.NewRequest("PUT", "/tasks/"+taskMap["id"].(string), bytes.NewBuffer(reqBody))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Testing GET
	req, _ = http.NewRequest("GET", "/tasks", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	body = make(map[string]any)
	err = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	taskMap = body["result"].([]any)[0].(map[string]any)
	assert.Equal(t, "Testing", taskMap["name"])
	assert.Equal(t, "1", taskMap["status"])
	assert.Len(t, body, 1)

	// Testing DELETE
	req, _ = http.NewRequest("DELETE", "/tasks/"+taskMap["id"].(string), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Testing GET
	req, _ = http.NewRequest("GET", "/tasks", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	body = make(map[string]any)
	err = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Len(t, body["result"], 0)
}
