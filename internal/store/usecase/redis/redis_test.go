package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	storage := NewRedisStorage()

	// testing database connection
	pong, err := storage.Client.Ping(context.TODO()).Result()
	assert.Equal(t, "PONG", pong)
	assert.Equal(t, nil, err)

	// testing CRUD functions
	// Create
	err = storage.Create(map[string]any{"name": "testing", "status": 0})
	assert.Equal(t, nil, err)

	// Read
	val := storage.Read()
	v := val[0]
	assert.Equal(t, "testing", v["name"])
	assert.Equal(t, "0", v["status"])

	// Update
	err = storage.Update(val[0]["id"], map[string]any{"name": "testing", "status": 1})
	assert.Equal(t, nil, err)

	// Read updated
	val = storage.Read()
	v = val[0]
	assert.Equal(t, "testing", v["name"])
	assert.Equal(t, "1", v["status"])

	// Delete
	err = storage.Delete(val[0]["id"])
	assert.Equal(t, nil, err)

	// Read delete
	val = storage.Read()
	assert.Equal(t, 0, len(val))
}
