package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
)

type CardUpdateHandler struct {
	api         *operations.HearthstoneCardServiceAPI
	cardService *application.CardService
}

func NewCardUpdateHandler(api *operations.HearthstoneCardServiceAPI, cardService *application.CardService) *CardUpdateHandler {
	return &CardUpdateHandler{
		api:         api,
		cardService: cardService,
	}
}

func (c *CardUpdateHandler) SetupHandler() {
	c.api.CardsPostCardsUpdateHandler = cards.PostCardsUpdateHandlerFunc(
		func(pup cards.PostCardsUpdateParams) middleware.Responder {
			log.Println("Handling request POST /cards/update...")
			defer log.Printf("Handled %s request", pup.HTTPRequest.URL)

			err := c.cardService.UpdateCards()
			if err != nil {
				log.Printf("Error occurred in POST /cards/update: %v", err)
				errorMessage := utils.CreateErrorMessage(500, "Something went wrong with updating cards")
				return cards.NewGetCardsInternalServerError().WithPayload(errorMessage)
			}

			return cards.NewPostCardsUpdateOK()
		},
	)
}
