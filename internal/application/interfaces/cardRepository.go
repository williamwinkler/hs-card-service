package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type CardRepository interface {
	FindAll(ctx context.Context) ([]domain.Card, error)
	FindWithFilter(ctx context.Context, filter domain.CardFilter, page int, limit int) ([]domain.Card, error)
	FindRichWithFilter(ctx context.Context, filter domain.CardFilter, page int, limit int) ([]domain.RichCard, error)
	InsertOne(ctx context.Context, card domain.Card) error
	InsertMany(ctx context.Context, cards []domain.Card) error
	UpdateOne(ctx context.Context, card domain.Card) error
	DeleteOne(ctx context.Context, card domain.Card) error
	DeleteAll(ctx context.Context) error
	CountWithFilter(ctx context.Context, filter domain.CardFilter) (int64, error)
}
