package handlers

import (
	"log"
	"math"

	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"go.mongodb.org/mongo-driver/bson"
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
	c.api.CardsGetRichcardsHandler = cards.GetRichcardsHandlerFunc(
		func(params cards.GetRichcardsParams) middleware.Responder {

			filter := bson.M{}
			if params.Name != nil {
				filter["name"] = bson.M{"$regex": ".*" + *params.Name + ".*", "$options": "i"}
			}
			if params.ManaCost != nil {
				// if manacost is 99 -> fetch cards with manacost greater than 7
				if *params.ManaCost == 99 {
					filter["manacost"] = bson.M{"$gte": 7}
				} else {
					filter["manacost"] = params.ManaCost
				}
			}
			if params.Health != nil {
				filter["health"] = params.Health
			}
			if params.Attack != nil {
				filter["attack"] = params.Attack
			}
			if params.Class != nil {
				filter["classid"] = params.Class
			}
			if params.Rarity != nil {
				filter["rarityid"] = params.Rarity
			}
			if params.Type != nil {
				filter["cardtypeid"] = params.Type
			}
			if params.Set != nil {
				filter["cardsetid"] = params.Set
			}
			if params.Keywords != nil && len(params.Keywords) > 0 {
				filter["keywordids"] = bson.M{"$all": params.Keywords}
			}

			foundCards, count, err := c.cardService.GetRichCards(filter, int(*params.Page), int(*params.Limit))
			if err != nil {
				errorMessage := utils.CreateErrorMessage(500, "Somthing went wrong with getting rich cards")
				return cards.NewGetRichcardsInternalServerError().WithPayload(errorMessage)
			}

			mappedCards := mapRichCardsToExternal(foundCards)
			pageCount := math.Ceil(float64(count) / float64(*params.Limit))

			log.Printf("Handled %s request (%d)", params.HTTPRequest.URL, len(mappedCards))
			response := models.RichCards{
				Page:      *params.Page,
				PageCount: int64(pageCount),
				CardCount: int64(count),
				Cards:     mappedCards,
			}
			return cards.NewGetRichcardsOK().WithPayload(&response)
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
