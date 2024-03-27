package items

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnify(t *testing.T) {
	task := Task{Name: "test", Status: 0}
	task.Unify()
	assert.Equal(t, Task{Name: "test", Status: 0}, task)
	task = Task{Name: "test", Status: 1}
	task.Unify()
	assert.Equal(t, Task{Name: "test", Status: 1}, task)
	task = Task{Name: "test", Status: 3}
	task.Unify()
	assert.Equal(t, Task{Name: "test", Status: 1}, task)
	task = Task{Name: "test", Status: -5}
	task.Unify()
	assert.Equal(t, Task{Name: "test", Status: 0}, task)
}
