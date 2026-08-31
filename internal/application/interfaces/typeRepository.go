package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type TypeRepository interface {
	InsertMany(ctx context.Context, types []domain.Type) error
	DeleteAll(ctx context.Context) error
}
