package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type RarityRepository interface {
	InsertMany(ctx context.Context, rarities []domain.Rarity) error
	DeleteAll(ctx context.Context) error
}
