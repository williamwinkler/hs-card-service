// This mock is manually maintained

package mocks

import "github.com/williamwinkler/hs-card-service/internal/domain"

type CardRepository struct {
	Cards map[int]domain.Card
}

func NewCardRepository() CardRepository {
	return CardRepository{
		Cards: make(map[int]domain.Card),
	}
}

func (c *CardRepository) InsertOne(card domain.Card) error {
	return nil
}

func (c *CardRepository) InsertMany(cards []domain.Card) error {
	for _, card := range cards {
		c.Cards[card.ID] = card
	}

	return nil
}

func (c *CardRepository) FindAll() ([]domain.Card, error) {
	var cards []domain.Card
	for _, card := range c.Cards {
		cards = append(cards, card)
	}

	return cards, nil
}

func (c *CardRepository) UpdateOne(card domain.Card) error {
	c.Cards[card.ID] = card
	return nil
}

func (c *CardRepository) DeleteOne(card domain.Card) error {
	delete(c.Cards, card.ID)
	return nil
}

func (c *CardRepository) DeleteAll() error {
	return nil
}
