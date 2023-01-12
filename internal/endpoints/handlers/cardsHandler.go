package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/domain"
)

type CardHandler struct {
	api         *operations.HearthstoneCardServiceAPI
	cardService *application.CardService
}

func NewCardHandler(api *operations.HearthstoneCardServiceAPI, cardService *application.CardService) *CardHandler {
	return &CardHandler{
		api:         api,
		cardService: cardService,
	}
}

func (c *CardHandler) SetupHandler() {
	c.api.CardsGetCardsHandler = cards.GetCardsHandlerFunc(
		func(gcp cards.GetCardsParams) middleware.Responder {

			foundCards, err := c.cardService.GetCards(gcp)
			if err != nil {
				return cards.NewGetCardsInternalServerError()
			}

			mappedCards := mapDomainCardsToExternal(foundCards)

			log.Printf("Handled %s request (%d)", gcp.HTTPRequest.URL, len(mappedCards))
			return cards.NewGetCardsOK().WithPayload(mappedCards)
		},
	)
}

func mapDomainCardsToExternal(cards []domain.Card) []*models.Card {
	var mappedCards []*models.Card
	for _, card := range cards {
		var c models.Card
		c.ID = int64(card.ID)
		c.ArtistName = card.ArtistName
		c.Attack = int64(card.Attack)
		c.CardSetID = int64(card.CardSetID)
		c.CardTypeID = int64(card.CardTypeID)
		c.ClassID = int64(card.ClassID)
		c.Collectible = int64(card.Collectible)
		c.FlavorText = card.FlavorText
		c.Health = int64(card.Health)
		c.Image = card.Image
		c.ImageGold = card.ImageGold
		c.ManaCost = int64(card.ManaCost)
		c.Name = card.Name
		c.ParentID = int64(card.ParentID)
		c.RarityID = int64(card.RarityID)
		c.Text = card.Text
		c.Duals = &models.Duals{
			Constructed: card.Duels.Constructed,
			Relevant:    card.Duels.Relevant,
		}

		mappedCards = append(mappedCards, &c)
	}

	return mappedCards
}
