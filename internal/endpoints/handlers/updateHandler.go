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
				c.UpdateWithRetries(ctx, 3, time.Second)
			}()

			return update.NewPostUpdateAccepted()
		})
}

func (c *CardUpdateHandler) UpdateWithRetries(ctx context.Context, maxRetries int, retryDelay time.Duration) {
	logging.Infof(ctx, "Handling request POST /cards/update...")
	defer logging.Debugf(ctx, "Handled /update request")

	retryFunc := func(updateFunc func() error, serviceName string) error {
		retries := 0
		for retries < maxRetries {
			err := updateFunc()
			if err == nil {
				return nil
			}
			logging.Errorf(ctx, "Error occurred in POST /cards/update (%s): %v", serviceName, err)
			retries++
			time.Sleep(retryDelay)
		}
		return fmt.Errorf("maximum retries reached for %s", serviceName)
	}

	type updateJob struct {
		name string
		fn   func() error
	}

	metadataJobs := []updateJob{
		{name: "set", fn: func() error { return c.setService.Update(ctx) }},
		{name: "class", fn: func() error { return c.classService.Update(ctx) }},
		{name: "rarity", fn: func() error { return c.rarityService.Update(ctx) }},
		{name: "type", fn: func() error { return c.typeService.Update(ctx) }},
		{name: "keyword", fn: func() error { return c.keywordService.Update(ctx) }},
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

	for err := range metadataErrors {
		logging.Errorf(ctx, "Aborting card update because metadata update failed: %v", err)
		return
	}

	if err := retryFunc(func() error { return c.cardService.Update(ctx) }, "card"); err != nil {
		logging.Errorf(ctx, "Card update failed: %v", err)
	}
}
