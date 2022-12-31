package main

import (
	"fmt"
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

	cards, err := hsClient.GetCardsWithPagination(1, 20)
	if err != nil {
		log.Fatal(err)
	}
	for _, card := range cards {
		fmt.Println(card.Name)
	}
}
