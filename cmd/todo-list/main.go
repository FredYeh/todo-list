package main

import (
	"log"
	"os"

	"github.com/FredYeh/todo-list/internal/router"
	"github.com/spf13/viper"
)

func main() {
	var config string
	if argvLen := len(os.Args); argvLen > 1 {
		config = os.Args[1]
	}
	app := router.Router(config)

	svr := viper.GetString("application.address") + ":" + viper.GetString("application.port")
	if err := app.Run(svr); err != nil {
		log.Fatal(err)
	}
}
