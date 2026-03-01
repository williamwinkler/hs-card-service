package handlers

import (
	"math"

	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/logging"
)

type RichCardHandler struct {
	api         *operations.HearthstoneCardServiceAPI
	cardService *application.CardService
}

func NewRichCardHandler(api *operations.HearthstoneCardServiceAPI, cardService *application.CardService) *RichCardHandler {
	return &RichCardHandler{
		api:         api,
		cardService: cardService,
	}
}

func (c *RichCardHandler) SetupHandler() {
	c.api.CardsGetRichCardsHandler = cards.GetRichCardsHandlerFunc(
		func(params cards.GetRichCardsParams) middleware.Responder {
			ctx := params.HTTPRequest.Context()

			filter := newCardFilter(params.Name, params.ManaCost, params.Health, params.Attack, params.Class, params.Rarity, params.Type, params.Set, params.Keywords)

			foundCards, count, err := c.cardService.GetRichCards(filter, int(*params.Page), int(*params.Limit))
			if err != nil {
				errorMessage := utils.CreateErrorMessage(500, "Somthing went wrong with getting rich cards")
				return cards.NewGetRichCardsInternalServerError().WithPayload(errorMessage)
			}

			mappedCards := mapRichCardsToExternal(foundCards)
			pageCount := math.Ceil(float64(count) / float64(*params.Limit))

			logging.Debugf(ctx, "Handled %s request (%d)", params.HTTPRequest.URL, len(mappedCards))
			response := models.RichCards{
				Page:      *params.Page,
				PageCount: int64(pageCount),
				CardCount: int64(count),
				Cards:     mappedCards,
			}
			return cards.NewGetRichCardsOK().WithPayload(&response)
		},
	)
}

func mapRichCardsToExternal(cards []domain.RichCard) []*models.RichCard {
	var mappedCards []*models.RichCard
	for _, card := range cards {
		var c models.RichCard
		c.ID = int64(card.ID)
		c.ArtistName = card.ArtistName
		c.Attack = int64(card.Attack)
		c.CardSet = card.CardSet
		c.CardType = card.CardType
		c.Class = card.Class
		c.Collectible = int64(card.Collectible)
		c.FlavorText = card.FlavorText
		c.Health = int64(card.Health)
		c.Image = card.Image
		c.ImageGold = card.ImageGold
		c.ManaCost = int64(card.ManaCost)
		c.Name = card.Name
		c.ParentID = int64(card.ParentID)
		c.Rarity = card.Rarity
		c.Text = card.Text
		c.Keywords = card.Keywords
		c.Duals = &models.Duals{
			Constructed: card.Duels.Constructed,
			Relevant:    card.Duels.Relevant,
		}

		mappedCards = append(mappedCards, &c)
	}

	return mappedCards
}
