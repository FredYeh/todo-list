package main

import (
	"log"
	"os"

	"github.com/FredYeh/todo-list/internal/router"
)

func main() {
	var config string
	if argvLen := len(os.Args); argvLen > 1 {
		config = os.Args[1]
	}
	app := router.Router(config)

	if err := app.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
