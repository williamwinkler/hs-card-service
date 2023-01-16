package interfaces

import "github.com/williamwinkler/hs-card-service/internal/domain"

type RarityRepository interface {
	InsertMany(raritys []domain.Rarity) error
	DeleteAll() error
}
