package main

import (
	"github.com/rs/zerolog/log"

	"github.com/guilhermevicente/person-management/api"
)

func main() {
	server := api.NewServer()
	server.ConfigRoutes()
	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msgf("Failed to start app")
	}
}
