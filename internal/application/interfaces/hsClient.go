package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type HsClient interface {
	GetCardsWithPagination(ctx context.Context, page int, pageSize int) ([]domain.Card, error)
	GetAllCards(ctx context.Context) ([]domain.Card, error)
	GetSets(ctx context.Context) ([]domain.Set, error)
	GetClasses(ctx context.Context) ([]domain.Class, error)
	GetRarities(ctx context.Context) ([]domain.Rarity, error)
	GetTypes(ctx context.Context) ([]domain.Type, error)
	GetKeywords(ctx context.Context) ([]domain.Keyword, error)
}
