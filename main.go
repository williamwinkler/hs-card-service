package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/clients"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	hsClient, err := clients.NewHsClient()
	if err != nil {
		log.Fatalf("Failed to start HsClient: %v", err)
	}

	cardRepository, err := repositories.NewCardRepository()
	if err != nil {
		log.Fatalf("Failed to start cardRepository: %v", err)
	}

	cardService := application.NewCardService(hsClient, cardRepository)
	cardService.UpdateCards()

	log.Println("Program Ended")
}
