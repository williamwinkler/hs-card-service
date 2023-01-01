package interfaces

import "github.com/williamwinkler/hs-card-service/domain"

type HsClient interface {
	GetCardsWithPagination(page int, pageSize int) ([]domain.Card, error)
	GetAllCards() ([]domain.Card, error)
}
