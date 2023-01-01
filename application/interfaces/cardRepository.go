package interfaces

import (
	"context"

	"github.com/williamwinkler/hs-card-service/domain"
)

type CardRepository interface {
	Insert(context.Context, domain.Card) error
	InsertMany(cards []domain.Card) error
}
