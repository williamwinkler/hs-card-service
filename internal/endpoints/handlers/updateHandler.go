package handlers

import (
	"log"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/update"
	"github.com/williamwinkler/hs-card-service/internal/application"
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
		func(pup update.PostUpdateParams) middleware.Responder {

			go c.UpdateWithRetries(3, 1000*time.Millisecond)

			return update.NewPostUpdateOK()
		},
	)
}

func (c *CardUpdateHandler) UpdateWithRetries(maxRetries int, retryDelay time.Duration) {
	log.Println("Handling request POST /cards/update...")
	defer log.Printf("Handled /update request")

	retryFunc := func(updateFunc func() error, serviceName string) {
		retries := 0
		for retries < maxRetries {
			err := updateFunc()
			if err == nil {
				return // Success, exit the retry loop
			}
			log.Printf("Error occurred in POST /cards/update (%s): %v", serviceName, err)
			retries++
			time.Sleep(retryDelay)
		}
		log.Printf("Maximum retries reached for %s. Giving up.", serviceName)
	}

	retryFunc(c.cardService.Update, "card")
	retryFunc(c.setService.Update, "set")
	retryFunc(c.classService.Update, "class")
	retryFunc(c.rarityService.Update, "rarity")
	retryFunc(c.typeService.Update, "type")
	retryFunc(c.keywordService.Update, "keyword")
}
