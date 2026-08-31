package application

import (
	"context"
	"time"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/observability"
)

type CardService struct {
	hsClient     interfaces.HsClient
	cardRepo     interfaces.CardRepository
	cardMetaRepo interfaces.UpdateMetaRepository
}

func NewCardService(hsclient interfaces.HsClient, cardRepo interfaces.CardRepository, cardMetaRepo interfaces.UpdateMetaRepository) *CardService {
	return &CardService{
		hsClient:     hsclient,
		cardRepo:     cardRepo,
		cardMetaRepo: cardMetaRepo,
	}
}

func (c *CardService) GetCards(ctx context.Context, filter domain.CardFilter, page int, limit int) ([]domain.Card, int64, error) {
	ctx, span := observability.Start(ctx, "cards.get")
	defer span.End()

	cards, err := c.cardRepo.FindWithFilter(ctx, filter, page, limit)
	if err != nil {
		observability.Fail(span, "database_failed")
		return nil, 0, err
	}

	count, err := c.cardRepo.CountWithFilter(ctx, filter)
	if err != nil {
		observability.Fail(span, "database_failed")
		return nil, 0, err
	}

	return cards, count, nil
}

func (c *CardService) GetRichCards(ctx context.Context, filter domain.CardFilter, page int, limit int) ([]domain.RichCard, int64, error) {
	ctx, span := observability.Start(ctx, "cards.get_rich")
	defer span.End()

	richCards, err := c.cardRepo.FindRichWithFilter(ctx, filter, page, limit)
	if err != nil {
		observability.Fail(span, "database_failed")
		return nil, 0, err
	}

	count, err := c.cardRepo.CountWithFilter(ctx, filter)
	if err != nil {
		observability.Fail(span, "database_failed")
		return nil, 0, err
	}

	return richCards, count, nil
}

func (c *CardService) Update(ctx context.Context) error {
	ctx, span := observability.Start(ctx, "cards.update.cards")
	defer span.End()

	cards, err := c.hsClient.GetAllCards(ctx)
	if err != nil {
		observability.Fail(span, "upstream_failed")
		return err
	}

	if err := c.cardRepo.DeleteAll(ctx); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}

	if err := c.cardMetaRepo.InsertOne(ctx, domain.CardMeta{
		Updated:   time.Now(),
		IsChanged: false,
	}); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}

	if err := c.cardRepo.InsertMany(ctx, cards); err != nil {
		observability.Fail(span, "database_failed")
		return err
	}
	return nil
}
