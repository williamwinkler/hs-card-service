package interfaces

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type CardRepository interface {
	InsertOne(card domain.Card) error
	InsertMany(cards []domain.Card) error
	FindAll() ([]domain.Card, error)
	FindWithFilter(bson.M) ([]domain.Card, error)
	UpdateOne(domain.Card) error
	DeleteOne(domain.Card) error
	DeleteAll() error
}
