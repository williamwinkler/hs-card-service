package interfaces

import "github.com/williamwinkler/hs-card-service/internal/domain"

type SetRepository interface {
	InsertMany(sets []domain.Set) error
	DeleteAll() error
}
