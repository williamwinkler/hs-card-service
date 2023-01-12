package interfaces

import "github.com/williamwinkler/hs-card-service/internal/domain"

type UpdateMetaRepository interface {
	InsertOne(domain.CardMeta) error
	FindNewest() (domain.CardMeta, error)
}
