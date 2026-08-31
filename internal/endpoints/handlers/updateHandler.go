package handlers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/update"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/logging"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
)

type CardUpdateHandler struct {
	api            *operations.HearthstoneCardServiceAPI
	cardService    *application.CardService
	setService     *application.SetService
	classService   *application.ClassService
	rarityService  *application.RarityService
	typeService    *application.TypeService
	keywordService *application.KeywordService
}

func NewCardUpdateHandler(
	api *operations.HearthstoneCardServiceAPI,
	cardService *application.CardService,
	setService *application.SetService,
	classService *application.ClassService,
	rarityService *application.RarityService,
	typeService *application.TypeService,
	keywordsService *application.KeywordService,
) *CardUpdateHandler {
	return &CardUpdateHandler{
		api:            api,
		cardService:    cardService,
		setService:     setService,
		classService:   classService,
		rarityService:  rarityService,
		typeService:    typeService,
		keywordService: keywordsService,
	}
}

func (c *CardUpdateHandler) SetupHandler() {
	c.api.UpdatePostUpdateHandler = update.PostUpdateHandlerFunc(
		func(pup update.PostUpdateParams, i interface{}) middleware.Responder {
			ctx := context.Background()
			if pup.HTTPRequest != nil {
				// The response is accepted immediately, so preserve trace values while
				// decoupling the background update from request cancellation.
				ctx = context.WithoutCancel(pup.HTTPRequest.Context())
			}
			ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)

			go func() {
				defer cancel()
				workflowCtx, span := observability.Start(ctx, "cards.update")
				defer span.End()
				if err := c.UpdateWithRetries(workflowCtx, 3, time.Second); err != nil {
					observability.Fail(span, "update_failed")
				}
			}()

			return update.NewPostUpdateAccepted()
		})
}

func (c *CardUpdateHandler) UpdateWithRetries(ctx context.Context, maxRetries int, retryDelay time.Duration) error {
	logging.Infof(ctx, "Handling request POST /cards/update")
	defer logging.Debugf(ctx, "Handled POST /cards/update")

	retryFunc := func(updateFunc func(context.Context) error, resource string) error {
		for attempt := 1; attempt <= maxRetries; attempt++ {
			attemptCtx, span := observability.Start(ctx, "cards.update.attempt",
				attribute.String("app.update.resource", resource),
				attribute.Int("app.retry.attempt", attempt),
			)
			err := updateFunc(attemptCtx)
			if err == nil {
				span.End()
				return nil
			}
			observability.Fail(span, "update_attempt_failed")
			span.End()
			logging.Errorf(attemptCtx, "POST /cards/update attempt failed for %s", resource)
			if err := waitForRetry(attemptCtx, retryDelay); err != nil {
				return err
			}
		}
		return fmt.Errorf("maximum retries reached for %s", resource)
	}

	type updateJob struct {
		name string
		fn   func(context.Context) error
	}

	metadataJobs := []updateJob{
		{name: "set", fn: c.setService.Update},
		{name: "class", fn: c.classService.Update},
		{name: "rarity", fn: c.rarityService.Update},
		{name: "type", fn: c.typeService.Update},
		{name: "keyword", fn: c.keywordService.Update},
	}

	metadataErrors := make(chan error, len(metadataJobs))
	var wg sync.WaitGroup
	for _, job := range metadataJobs {
		job := job
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := retryFunc(job.fn, job.name); err != nil {
				metadataErrors <- err
			}
		}()
	}
	wg.Wait()
	close(metadataErrors)

	for range metadataErrors {
		logging.Errorf(ctx, "Aborting card update because metadata update failed")
		return fmt.Errorf("metadata update failed")
	}

	if err := retryFunc(c.cardService.Update, "card"); err != nil {
		logging.Errorf(ctx, "Card update failed")
		return err
	}
	return nil
}

func waitForRetry(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
