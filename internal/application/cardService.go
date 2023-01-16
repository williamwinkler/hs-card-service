package application

import (
	"time"

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

func (c *CardService) GetCards(filter bson.M, page int, limit int) ([]domain.Card, int64, error) {
	cards, err := c.cardRepo.FindWithFilter(filter, page, limit)
	if err != nil {
		return []domain.Card{}, 0, err
	}

	count, err := c.cardRepo.CountWithFilter(filter)
	if err != nil {
		return []domain.Card{}, 0, err
	}

	return cards, count, nil
}

func (c *CardService) GetRichCards(filter bson.M, page int, limit int) ([]domain.RichCard, int64, error) {
	richCards, err := c.cardRepo.FindRichWithFilter(filter, page, limit)
	if err != nil {
		return []domain.RichCard{}, 0, err
	}

	count, err := c.cardRepo.CountWithFilter(filter)
	if err != nil {
		return []domain.RichCard{}, 0, err
	}

	return richCards, count, nil
}

func (c *CardService) Update() error {
	// TODO: make it smarter, so it only deletes/updates/adds affected cards
	cards, err := c.hsClient.GetAllCards()
	if err != nil {
		return err
	}

	err = c.cardRepo.DeleteAll()
	if err != nil {
		return err
	}

	// TODO: make sure to save the if the update made changes
	// currently it only saves that a update happend, but not what happened
	c.cardMetaRepo.InsertOne(domain.CardMeta{
		Updated:   time.Now(),
		IsChanged: false,
	})

	return c.cardRepo.InsertMany(cards)
}
