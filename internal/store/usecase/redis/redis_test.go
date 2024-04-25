package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	config := redisConfig{Host: "localhost", Port: "6379"}
	storage := Storage{Client: NewClient(&config)}

	// testing database connection
	pong, err := storage.Client.Ping(context.TODO()).Result()
	assert.Equal(t, "PONG", pong)
	assert.Equal(t, nil, err)

	// testing CRUD functions
	// Create
	id, err := storage.Create(map[string]any{"name": "testing", "status": 0})
	assert.Equal(t, nil, err)

	// Read
	val, _ := storage.Read()
	v := val[0]
	assert.Equal(t, "testing", v["name"])
	assert.Equal(t, "0", v["status"])

	// Update
	err = storage.Update(id, map[string]any{"name": "testing", "status": 1})
	assert.Equal(t, nil, err)

	// Read updated
	val, _ = storage.Read()
	v = val[0]
	assert.Equal(t, "testing", v["name"])
	assert.Equal(t, "1", v["status"])

	// Delete
	err = storage.Delete(id)
	assert.Equal(t, nil, err)

	// Read delete
	val, _ = storage.Read()
	assert.Equal(t, 0, len(val))
}
