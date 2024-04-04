package handlers

import (
	"log"
	"net/http"

	"github.com/FredYeh/todo-list/internal/store"
	"github.com/FredYeh/todo-list/internal/store/items"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Storage store.TaskStorage
}

func (h *TaskHandler) GetHandler(c *gin.Context) {
	res := h.Storage.Read()
	c.JSON(http.StatusOK, res)
}

func (h *TaskHandler) PostHandler(c *gin.Context) {
	task := new(items.Task)
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		if err := h.Storage.Create(task); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusCreated, "New task created.")
		}
	}
}

func (h *TaskHandler) PutHandler(c *gin.Context) {
	id := c.Param("id")
	task := new(items.Task)
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		if err := h.Storage.Update(id, task); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusOK, "Task updated.")
		}
	}
}

func (h *TaskHandler) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.Storage.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, "Task deleted.")
	}
}
