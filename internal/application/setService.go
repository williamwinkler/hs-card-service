package application

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
)

type SetService struct {
	setRepo  interfaces.SetRepository
	hsClient interfaces.HsClient
}

func NewSetService(setRepo interfaces.SetRepository, hsClient interfaces.HsClient) *SetService {
	return &SetService{
		setRepo:  setRepo,
		hsClient: hsClient,
	}
}

func (c *SetService) Update(ctx context.Context) error {
	ctx, span := observability.Start(ctx, "cards.update.metadata", attribute.String("app.update.resource", "set"))
	defer span.End()

	sets, err := c.hsClient.GetSets(ctx)
	if err != nil {
		observability.Fail(span, "upstream_failed")
		return err
	}
	if err := c.setRepo.DeleteAll(ctx); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	if err := c.setRepo.InsertMany(ctx, sets); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	return nil
}
