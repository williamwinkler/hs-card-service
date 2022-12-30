package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/infrastructure/clients"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	hsClient, err := clients.NewHsClient()
	if err != nil {
		log.Fatal(err)
	}

	cards, err := hsClient.GetCardsWithPagination(1, 1)
	if err != nil {
		log.Fatal(err)
	}
	println(cards)
}
