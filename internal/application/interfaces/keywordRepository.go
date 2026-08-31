package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type KeywordRepository interface {
	InsertMany(ctx context.Context, keywords []domain.Keyword) error
	DeleteAll(ctx context.Context) error
}
