package application

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
)

type KeywordService struct {
	keywordRepo interfaces.KeywordRepository
	hsClient    interfaces.HsClient
}

func NewKeywordService(keywordRepo interfaces.KeywordRepository, hsClient interfaces.HsClient) *KeywordService {
	return &KeywordService{
		keywordRepo: keywordRepo,
		hsClient:    hsClient,
	}
}

func (c *KeywordService) Update(ctx context.Context) error {
	ctx, span := observability.Start(ctx, "cards.update.metadata", attribute.String("app.update.resource", "keyword"))
	defer span.End()

	keywords, err := c.hsClient.GetKeywords(ctx)
	if err != nil {
		observability.Fail(span, "upstream_failed")
		return err
	}
	if err := c.keywordRepo.DeleteAll(ctx); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	if err := c.keywordRepo.InsertMany(ctx, keywords); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	return nil
}
