package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/update"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
)

type CardUpdateHandler struct {
	api         *operations.HearthstoneCardServiceAPI
	cardService *application.CardService
	setService  *application.SetService
}

func NewCardUpdateHandler(api *operations.HearthstoneCardServiceAPI, cardService *application.CardService, setService *application.SetService) *CardUpdateHandler {
	return &CardUpdateHandler{
		api:         api,
		cardService: cardService,
		setService:  setService,
	}
}

func (c *CardUpdateHandler) SetupHandler() {
	c.api.UpdatePostUpdateHandler = update.PostUpdateHandlerFunc(
		func(pup update.PostUpdateParams) middleware.Responder {
			log.Println("Handling request POST /cards/update...")
			defer log.Printf("Handled %s request", pup.HTTPRequest.URL)

			err := c.cardService.Update()
			if err != nil {
				log.Printf("Error occurred in POST /cards/update: %v", err)
				errorMessage := utils.CreateErrorMessage(500, "Something went wrong with updating cards")
				return cards.NewGetCardsInternalServerError().WithPayload(errorMessage)
			}

			err = c.setService.Update()
			if err != nil {
				log.Printf("Error occurred in POST /cards/update: %v", err)
				errorMessage := utils.CreateErrorMessage(500, "Something went wrong with updating sets")
				return cards.NewGetCardsInternalServerError().WithPayload(errorMessage)
			}

			return update.NewPostUpdateOK()
		},
	)
}
