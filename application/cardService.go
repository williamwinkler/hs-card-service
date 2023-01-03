package application

import (
	"fmt"

	"github.com/williamwinkler/hs-card-service/application/interfaces"
)

type CardService struct {
	HsClient interfaces.HsClient
	CardRepo interfaces.CardRepository
}

func NewCardService(hsclient interfaces.HsClient, cardRepo interfaces.CardRepository) *CardService {
	return &CardService{
		HsClient: hsclient,
		CardRepo: cardRepo,
	}
}

func (c *CardService) UpdateCards() error {
	// newCards, err := c.HsClient.GetAllCards()
	// if err != nil {
	// 	return err
	// }

	oldCards, err := c.CardRepo.FindAll()
	if err != nil {
		return err
	}

	for i, card := range oldCards {
		fmt.Printf("%d %s\n", i+1, card.Name)
	}

	return nil
}
