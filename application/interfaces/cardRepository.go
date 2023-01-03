package interfaces

import (
	"github.com/williamwinkler/hs-card-service/domain"
)

type CardRepository interface {
	Insert(card domain.Card) error
	InsertMany(cards []domain.Card) error
	FindAll() ([]domain.Card, error)
}
