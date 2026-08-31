package application

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
)

type ClassService struct {
	classRepo interfaces.ClassRepository
	hsClient  interfaces.HsClient
}

func NewClassService(ClassRepo interfaces.ClassRepository, hsClient interfaces.HsClient) *ClassService {
	return &ClassService{
		classRepo: ClassRepo,
		hsClient:  hsClient,
	}
}

func (c *ClassService) Update(ctx context.Context) error {
	ctx, span := observability.Start(ctx, "cards.update.metadata", attribute.String("app.update.resource", "class"))
	defer span.End()

	classes, err := c.hsClient.GetClasses(ctx)
	if err != nil {
		observability.Fail(span, "upstream_failed")
		return err
	}
	if err := c.classRepo.DeleteAll(ctx); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	if err := c.classRepo.InsertMany(ctx, classes); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	return nil
}
