package router

import (
	"log"

	"github.com/FredYeh/todo-list/internal/handlers"
	"github.com/FredYeh/todo-list/internal/store/usecase/redis"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setEnv(config string) error {
	if config == "" {
		config = "test"
	}
	viper.AddConfigPath("./config")
	viper.SetConfigName(config)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func Router(config string) *gin.Engine {
	if err := setEnv(config); err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	gin.SetMode(viper.GetString("mode"))

	storage := redis.NewRedisStorage()

	taskapi := router.Group("/tasks")
	thandler := handlers.TaskHandler{Storage: storage}
	{
		taskapi.GET("/", thandler.GetHandler)
		taskapi.POST("/", thandler.PostHandler)
		taskapi.PUT("/:id", thandler.PutHandler)
		taskapi.DELETE("/:id", thandler.DeleteHandler)
	}

	return router
}
