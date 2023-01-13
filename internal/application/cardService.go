package application

import (
	"log"
	"sort"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type CardService struct {
	hsClient     interfaces.HsClient
	cardRepo     interfaces.CardRepository
	cardMetaRepo interfaces.UpdateMetaRepository
}

func NewCardService(hsclient interfaces.HsClient, cardRepo interfaces.CardRepository, cardMetaRepo interfaces.UpdateMetaRepository) *CardService {
	return &CardService{
		hsClient:     hsclient,
		cardRepo:     cardRepo,
		cardMetaRepo: cardMetaRepo,
	}
}

func (c *CardService) GetCards(params cards.GetCardsParams) ([]domain.Card, int64, error) {
	filter := bson.M{}

	if params.Name != nil {
		filter["name"] = params.Name
	}
	if params.ManaCost != nil {
		filter["manacost"] = params.ManaCost
	}
	if params.Health != nil {
		filter["health"] = params.Health
	}
	if params.Attack != nil {
		filter["attack"] = params.Attack
	}
	if params.Class != nil {
		filter["classid"] = params.Class
	}
	if params.Rarity != nil {
		filter["rarityid"] = params.Rarity
	}

	cards, err := c.cardRepo.FindWithFilter(filter, int(*params.Page), int(*params.Limit))
	if err != nil {
		return []domain.Card{}, 0, err
	}

	count, err := c.cardRepo.CountWithFilter(filter)
	if err != nil {
		return []domain.Card{}, 0, err
	}

	return cards, count, nil
}

func (c *CardService) UpdateCards() error {
	cards, err := c.hsClient.GetAllCards()
	if err != nil {
		return err
	}
	newCards := sortCardsByAlphabet(cards)

	oldCards, err := c.cardRepo.FindAll()
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

	cardsMetaUpdate := domain.NewCardMeta()

	// Remove old cards from the collection which are not in the new list
	for i, card := range oldCards {
		if !newCardsMap[card.ID] {
			err := c.cardRepo.DeleteOne(card)
			if err != nil {
				log.Fatalf("Card %d failed to be deleted: %v", card.ID, err)
				continue
			}
			log.Printf("Card %d was deleted", card.ID)
			delete(oldCardsMap, card.ID)
			oldCards = append(oldCards[:i], oldCards[i+1:]...)
			cardsMetaUpdate.DeleteCard(card)
		}
	}

	// Add new or update cards to the collection
	for i, card := range newCards {
		if !oldCardsMap[card.ID] {
			err := c.cardRepo.InsertOne(card)
			if err != nil {
				log.Fatalf("Card %d failed to inserted: %v", card.ID, err)
				continue
			}
			log.Printf("Card %d was added", card.ID)
			cardsMetaUpdate.AddCard(card)
		} else {
			if !card.Equals(oldCards[i]) {
				err := c.cardRepo.UpdateOne(card)
				if err != nil {
					log.Fatalf("Card %d failed to be updated: %v", card.ID, err)
					continue
				}
				log.Printf("Card %d was updated", card.ID)
				cardsMetaUpdate.UpdateCard(oldCards[i], card)
			}
		}
	}

	err = c.cardMetaRepo.InsertOne(cardsMetaUpdate)
	if err != nil {
		return err
	}

	return nil
}

func sortCardsByAlphabet(cards []domain.Card) []domain.Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Name < cards[i].Name
	})
	return cards
}
