package interfaces

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type CardRepository interface {
	FindAll() ([]domain.Card, error)
	FindWithFilter(filter bson.M, page int, limit int) ([]domain.Card, error)
	FindRichWithFilter(filter bson.M, page int, limit int) ([]domain.RichCard, error)
	InsertOne(card domain.Card) error
	InsertMany(cards []domain.Card) error
	UpdateOne(domain.Card) error
	DeleteOne(domain.Card) error
	DeleteAll() error
	CountWithFilter(filter bson.M) (int64, error)
}
