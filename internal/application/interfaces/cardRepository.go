package interfaces

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type CardRepository interface {
	FindAll() ([]domain.Card, error)
	FindWithFilter(filter domain.CardFilter, page int, limit int) ([]domain.Card, error)
	FindRichWithFilter(filter domain.CardFilter, page int, limit int) ([]domain.RichCard, error)
	InsertOne(card domain.Card) error
	InsertMany(cards []domain.Card) error
	UpdateOne(domain.Card) error
	DeleteOne(domain.Card) error
	DeleteAll() error
	CountWithFilter(filter domain.CardFilter) (int64, error)
}
