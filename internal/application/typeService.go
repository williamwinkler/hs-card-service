package application

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
)

type TypeService struct {
	typeRepo interfaces.TypeRepository
	hsClient interfaces.HsClient
}

func NewTypeService(typeRepo interfaces.TypeRepository, hsClient interfaces.HsClient) *TypeService {
	return &TypeService{
		typeRepo: typeRepo,
		hsClient: hsClient,
	}
}

func (c *TypeService) Update(ctx context.Context) error {
	ctx, span := observability.Start(ctx, "cards.update.metadata", attribute.String("app.update.resource", "type"))
	defer span.End()

	types, err := c.hsClient.GetTypes(ctx)
	if err != nil {
		observability.Fail(span, "upstream_failed")
		return err
	}
	if err := c.typeRepo.DeleteAll(ctx); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	if err := c.typeRepo.InsertMany(ctx, types); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	return nil
}
