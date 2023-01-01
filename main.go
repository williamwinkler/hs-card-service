package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/domain"
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

	card := domain.Card{
		Name: "test",
	}

	cardRepo.Insert(card)


	// hsClient, err := clients.NewHsClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cards, err := hsClient.GetCardsWithPagination(1, 287)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, card := range cards {
	// 	fmt.Println(card.Name)
	// }

	log.Println("Program Ended")
}
