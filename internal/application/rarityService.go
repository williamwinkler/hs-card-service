package application

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
)

type RarityService struct {
	rarityRepo interfaces.RarityRepository
	hsClient   interfaces.HsClient
}

func NewRarityService(rarityRepo interfaces.RarityRepository, hsClient interfaces.HsClient) *RarityService {
	return &RarityService{
		rarityRepo: rarityRepo,
		hsClient:   hsClient,
	}
}

func (c *RarityService) Update(ctx context.Context) error {
	ctx, span := observability.Start(ctx, "cards.update.metadata", attribute.String("app.update.resource", "rarity"))
	defer span.End()

	rarities, err := c.hsClient.GetRarities(ctx)
	if err != nil {
		observability.Fail(span, "upstream_failed")
		return err
	}
	if err := c.rarityRepo.DeleteAll(ctx); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	if err := c.rarityRepo.InsertMany(ctx, rarities); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	return nil
}
