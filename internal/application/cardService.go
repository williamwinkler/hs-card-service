package application

import (
	"sort"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/domain"
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
	cards, err := c.HsClient.GetAllCards()
	if err != nil {
		return err
	}
	newCards := sortCardsByAlphabet(cards)

	oldCards, err := c.CardRepo.FindAll()
	if err != nil {
		return err
	}

	newCardsMap := make(map[*domain.Card]bool)
	for _, card := range newCards {
		newCardsMap[&card] = true
	}

	oldCardsMap := make(map[*domain.Card]bool)
	for _, card := range oldCards {
		oldCardsMap[&card] = true
	}

	for i, card := range oldCards {
		if !newCardsMap[&card] {
			c.CardRepo.DeleteOne(card)
			oldCards = append(oldCards[:i], oldCards[i+1:]...)
		} else {
			if !card.IsEqual(newCards[i]) {
				c.CardRepo.UpdateOne(newCards[i])
				oldCards[i] = newCards[i]
			}
			delete(newCardsMap, &card)
			delete(oldCardsMap, &card)
		}
	}

	for card := range newCardsMap {
		c.CardRepo.InsertOne(*card)
	}

	return nil
}

func sortCardsByAlphabet(cards []domain.Card) []domain.Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Name < cards[i].Name
	})
	return cards
}
