package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type SetRepository interface {
	InsertMany(ctx context.Context, sets []domain.Set) error
	DeleteAll(ctx context.Context) error
}
