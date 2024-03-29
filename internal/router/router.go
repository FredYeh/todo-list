package router

import (
	"github.com/FredYeh/todo-list/internal/handlers"
	"github.com/FredYeh/todo-list/internal/store/usecase/redis"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetEnv(config string) {
	viper.AddConfigPath("./config")
	viper.SetConfigName(config)
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("mode", "debug")
		viper.Set("application.port", 80)
		viper.Set("database.host", "localhost")
		viper.Set("database.port", 6379)
	}
}

func Router(config string) *gin.Engine {
	SetEnv(config)
	router := gin.Default()
	gin.SetMode(viper.GetString("mode"))

	storage := redis.NewRedisStorage()
	taskapi := router.Group("/tasks")
	thandler := handlers.TaskHandler{Storage: storage}
	{
		taskapi.GET("", thandler.GetHandler)
		taskapi.POST("", thandler.PostHandler)
		taskapi.PUT("/:id", thandler.PutHandler)
		taskapi.DELETE("/:id", thandler.DeleteHandler)
	}

	return router
}
