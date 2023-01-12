package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/endpoints"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/clients"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	database, err := migrations.SetupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	hsClient, err := clients.NewHsClient()
	if err != nil {
		log.Fatalf("Failed to start HsClient: %v", err)
	}

	cardRepository := repositories.NewCardRepository(database.Db)
	cardMetaRepository := repositories.NewUpdateMetaRepository(database.Db)

	cardService := application.NewCardService(hsClient, cardRepository, cardMetaRepository)

	restServer := endpoints.NewRestServer(cardService, cardRepository)

	restServer.StartServer()

	log.Println("Program Ended")
}
