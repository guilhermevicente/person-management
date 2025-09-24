package main

import (
	"log"

	"github.com/guilhermevicente/person-management/api"
)

func main() {
	server := api.NewServer()
	server.ConfigRoutes()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
