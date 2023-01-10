package interfaces

import "github.com/williamwinkler/hs-card-service/internal/domain"

type CardMetaRepository interface {
	InsertOne(domain.CardMeta) error
	FindNewest() (domain.CardMeta, error)
}
