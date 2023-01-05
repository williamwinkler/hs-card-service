package interfaces

import (
	"github.com/williamwinkler/hs-card-service/domain"
)

type CardRepository interface {
	InsertOne(card domain.Card) error
	InsertMany(cards []domain.Card) error
	FindAll() ([]domain.Card, error)
	UpdateOne(domain.Card) error
	DeleteOne(domain.Card) error
	DeleteAll() error
}
