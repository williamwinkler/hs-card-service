package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type UpdateMetaRepository interface {
	InsertOne(ctx context.Context, cardMeta domain.CardMeta) error
	FindNewest(ctx context.Context) (domain.CardMeta, error)
}
