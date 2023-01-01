package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/infrastructure/clients"
	"github.com/williamwinkler/hs-card-service/infrastructure/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	cardRepo, err := repositories.NewCardRepository()
	if err != nil {
		log.Fatal(err)
	}

	hsClient, err := clients.NewHsClient()
	if err != nil {
		log.Fatal(err)
	}

	//cardService := application.NewCardService(hsClient, cardRepo)

	cards, err := hsClient.GetAllCards()

	log.Println("Inserting cards into database")
	err = cardRepo.InsertMany(cards)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Finished inserting cards into database")
	}

	log.Println("Program Ended")
}
