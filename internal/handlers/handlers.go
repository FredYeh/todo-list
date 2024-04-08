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
	c.JSON(http.StatusOK, gin.H{
		"msg":    "query successful",
		"result": res,
	})
}

func (h *TaskHandler) PostHandler(c *gin.Context) {
	task := items.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error: " + err.Error()})
	} else {
		if id, err := h.Storage.Create(task.Map()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "error: " + err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"msg": "New task created",
				"id":  id,
			})
		}
	}
}

func (h *TaskHandler) PutHandler(c *gin.Context) {
	id := c.Param("id")
	task := items.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error: " + err.Error()})
	} else {
		if err := h.Storage.Update(id, task.Map()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "error: " + err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Task updated.",
				"id":  id,
			})
		}
	}
}

func (h *TaskHandler) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.Storage.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error: " + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Task deleted.",
			"id":  id,
		})
	}
}
