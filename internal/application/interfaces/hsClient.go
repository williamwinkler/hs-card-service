package interfaces

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type HsClient interface {
	GetCardsWithPagination(page int, pageSize int) ([]domain.Card, error)
	GetAllCards() ([]domain.Card, error)
	GetSets() ([]domain.Set, error)
	GetClasses() ([]domain.Class, error)
	GetRarities() ([]domain.Rarity, error)
	GetTypes() ([]domain.Type, error)
	GetKeywords() ([]domain.Keyword, error)
}
