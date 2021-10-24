package main

import (
	"jwt-go/internal/server"
	"jwt-go/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".") // initialize config
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	server.Run(config) // Run server
}
