package application

import (
	"log"
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
	log.Println("Started updating cards...")

	cards, err := c.HsClient.GetAllCards()
	if err != nil {
		return err
	}
	newCards := sortCardsByAlphabet(cards)

	oldCards, err := c.CardRepo.FindAll()
	if err != nil {
		return err
	}

	newCardsMap := make(map[int]bool)
	for _, card := range newCards {
		newCardsMap[card.ID] = true
	}

	oldCardsMap := make(map[int]bool)
	for _, card := range oldCards {
		oldCardsMap[card.ID] = true
	}

	// Remove old cards from the collection which are not in the new list
	for i, card := range oldCards {
		if !newCardsMap[card.ID] {
			err := c.CardRepo.DeleteOne(card)
			if err != nil {
				log.Fatalf("Card %d failed to be deleted: %v", card.ID, err)
				continue
			}
			log.Printf("Card %d was deleted", card.ID)
			delete(oldCardsMap, card.ID)
			oldCards = append(oldCards[:i], oldCards[i+1:]...)
		}
	}

	// Add new or update cards to the collection
	for i, card := range newCards {
		if !oldCardsMap[card.ID] {
			err := c.CardRepo.InsertOne(card)
			if err != nil {
				log.Fatalf("Card %d failed to inserted: %v", card.ID, err)
				continue
			}
			log.Printf("Card %d was added", card.ID)
		} else {
			if !card.Equals(oldCards[i]) {
				err := c.CardRepo.UpdateOne(card)
				if err != nil {
					log.Fatalf("Card %d failed to be updated: %v", card.ID, err)
					continue
				}
				log.Printf("Card %d was updated", card.ID)
			}
		}
	}

	log.Println("Finished updating cards")
	return nil
}

func sortCardsByAlphabet(cards []domain.Card) []domain.Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Name < cards[i].Name
	})
	return cards
}
