package items

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	task := Task{Name: "test", Status: Todo}
	data, _ := json.Marshal(&task)
	expected, _ := json.Marshal(map[string]any{"name": "test", "status": "Todo"})
	assert.Equal(t, string(expected), string(data))

	task = Task{Name: "test", Status: Done}
	data, _ = json.Marshal(&task)
	expected, _ = json.Marshal(map[string]any{"name": "test", "status": "Done"})
	assert.Equal(t, string(expected), string(data))

	task = Task{Name: "test", Status: 3}
	data, err := json.Marshal(&task)
	assert.Equal(t, "", string(data))
	assert.Errorf(t, err, "3 is not a valid stutas")

	task = Task{}
	data = []byte(`{"name": "test", "status": "Todo"}`)
	err = json.Unmarshal(data, &task)
	assert.NoError(t, err)
	assert.Equal(t, Task{Name: "test", Status: Todo}, task)
	assert.Equal(t, 0, int(task.Status))

	task = Task{}
	data = []byte(`{"name": "test", "status": "Done"}`)
	err = json.Unmarshal(data, &task)
	assert.NoError(t, err)
	assert.Equal(t, Task{Name: "test", Status: Done}, task)
	assert.Equal(t, 1, int(task.Status))

	task = Task{}
	data = []byte(`{"name": "test", "status": "Hello"}`)
	err = json.Unmarshal(data, &task)
	assert.Errorf(t, err, "Hello is not a valid stataus")
}
