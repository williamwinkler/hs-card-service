package interfaces

import "github.com/williamwinkler/hs-card-service/internal/domain"

type KeywordRepository interface {
	InsertMany(keywords []domain.Keyword) error
	DeleteAll() error
}
