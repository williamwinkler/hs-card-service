package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/tests/mocks"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	newCards := []domain.Card{
		{
			ID:   1,
			Name: "card_a",
		},
		{
			ID:   2,
			Name: "card_b",
		},
		{
			ID:   4,
			Name: "card_d",
		},
	}

	cardRepo := mocks.NewCardRepository()

	cardRepo.InsertMany(newCards)

	cards, _ := cardRepo.FindAll()

	fmt.Println(cards)

	// cards, err := hsClient.GetAllCards()

	// log.Println("Inserting cards into database")
	// err = cardRepo.InsertMany(cards)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Finished inserting cards into database")
	// }

	log.Println("Program Ended")
}
