package interfaces

import "github.com/williamwinkler/hs-card-service/internal/domain"

type ClassRepository interface {
	InsertMany(classes []domain.Class) error
	DeleteAll() error
}
