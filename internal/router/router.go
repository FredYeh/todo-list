package router

import (
	"log"
	"net/http"

	"github.com/FredYeh/todo-list/internal/handlers"
	"github.com/FredYeh/todo-list/internal/store/usecase/redis"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetEnv(config string) {
	viper.AddConfigPath("./config")
	viper.SetConfigName(config)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Couldn't read config, using default config")
		viper.Set("mode", "debug")
		viper.Set("application.port", 80)
		viper.Set("database.host", "localhost")
		viper.Set("database.port", 6379)
	}
}

func Router(config string) *gin.Engine {
	SetEnv(config)
	gin.SetMode(viper.GetString("mode"))
	router := gin.Default()

	storage := redis.NewRedisStorage()
	taskapi := router.Group("/tasks")
	thandler := handlers.TaskHandler{Storage: storage}
	{
		taskapi.GET("", thandler.GetHandler)
		taskapi.POST("", thandler.PostHandler)
		taskapi.PUT("/:id", thandler.PutHandler)
		taskapi.DELETE("/:id", thandler.DeleteHandler)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 not found"})
	})

	return router
}
