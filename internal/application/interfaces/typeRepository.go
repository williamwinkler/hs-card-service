package interfaces

import "github.com/williamwinkler/hs-card-service/internal/domain"

type TypeRepository interface {
	InsertMany(types []domain.Type) error
	DeleteAll() error
}
