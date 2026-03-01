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
				ctx = pup.HTTPRequest.Context()
			}

			go c.UpdateWithRetries(ctx, 3, 1000*time.Millisecond)

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

	for err := range metadataErrors {
		logging.Errorf(ctx, "Aborting card update because metadata update failed: %v", err)
		return
	}

	if err := retryFunc(c.cardService.Update, "card"); err != nil {
		logging.Errorf(ctx, "Card update failed: %v", err)
	}
}
