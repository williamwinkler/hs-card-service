package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type ClassRepository interface {
	InsertMany(ctx context.Context, classes []domain.Class) error
	DeleteAll(ctx context.Context) error
}
