package application

import (
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
	newCards, err := c.HsClient.GetAllCards()
	if err != nil {
		return err
	}

	oldCards, err := c.CardRepo.FindAll()
	if err != nil {
		return err
	}

	if len(newCards) != len(oldCards) {
		// remove old
		// insert new
	}

	// otherwise go through card by card and check if there are any differences.
	// probably need to sort them even tho they come sorted the same from api and db

	return nil
}
