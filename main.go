package main

import (
	"github.com/rs/zerolog/log"

	"github.com/b4t3ou/csv-ingester/config"
	"github.com/b4t3ou/csv-ingester/storage"
	"github.com/b4t3ou/csv-ingester/storage/database"
)

const (
	// ServiceTypeServer representing the server service type
	ServiceTypeServer = "server"

	// ServiceTypeClient representing the client service type
	//ServiceTypeClient = "client"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Error().Err(err).Msg("failed to load config")
		return
	}

	switch cfg.ServiceType {
	case ServiceTypeServer:
		runServer(cfg)
	default:
		log.Fatal().Msg("service type not found")
	}
}

func runServer(cfg *config.Config) {
	server := storage.NewServer(cfg, database.NewMemory())
	server.ListenAndServe()
}
