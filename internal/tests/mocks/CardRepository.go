// This mock is manually maintained

package mocks

import (
	"context"
	"fmt"
	"reflect"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type CardRepository struct {
	Cards map[int]domain.Card // used like a InMemory db
}

func NewCardRepository() CardRepository {
	return CardRepository{
		Cards: make(map[int]domain.Card),
	}
}

func (c *CardRepository) InsertOne(ctx context.Context, card domain.Card) error {
	return nil
}

func (c *CardRepository) InsertMany(ctx context.Context, cards []domain.Card) error {
	for _, card := range cards {
		c.Cards[card.ID] = card
	}

	return nil
}

func (c *CardRepository) FindAll(ctx context.Context) ([]domain.Card, error) {
	var cards []domain.Card
	for _, card := range c.Cards {
		cards = append(cards, card)
	}

	return cards, nil
}

func (c *CardRepository) FindWithFilter(ctx context.Context, filter domain.CardFilter, page int, limit int) ([]domain.Card, error) {
	if !reflect.DeepEqual(filter, domain.CardFilter{}) {
		return []domain.Card{}, fmt.Errorf("Mock does not support filter")
	}
	return c.FindAll(ctx)
}

func (c *CardRepository) FindRichWithFilter(ctx context.Context, filter domain.CardFilter, page int, limit int) ([]domain.RichCard, error) {
	return []domain.RichCard{}, nil
}

func (c *CardRepository) UpdateOne(ctx context.Context, card domain.Card) error {
	c.Cards[card.ID] = card
	return nil
}

func (c *CardRepository) DeleteOne(ctx context.Context, card domain.Card) error {
	delete(c.Cards, card.ID)
	return nil
}

func (c *CardRepository) DeleteAll(ctx context.Context) error {
	return nil
}

func (c *CardRepository) CountWithFilter(ctx context.Context, filter domain.CardFilter) (int64, error) {
	return 0, nil
}
