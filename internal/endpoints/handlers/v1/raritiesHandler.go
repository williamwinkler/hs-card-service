package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/rarities"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

type RaritiesHandler struct {
	api        *operations.HearthstoneCardServiceAPI
	rarityRepo *repositories.RarityRepository
}

func NewRaritiesHandler(api *operations.HearthstoneCardServiceAPI, rarityRepo *repositories.RarityRepository) *RaritiesHandler {
	return &RaritiesHandler{
		api:        api,
		rarityRepo: rarityRepo,
	}
}

func (i *RaritiesHandler) SetupHandler() {
	i.api.RaritiesGetV1RaritiesHandler = rarities.GetV1RaritiesHandlerFunc(
		func(req rarities.GetV1RaritiesParams) middleware.Responder {
			defer log.Printf("Handled %s request", req.HTTPRequest.URL)

			cardRarities, err := i.rarityRepo.FindAll()
			if err != nil {
				log.Printf("Error occurred in GET /Rarities: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				rarities.NewGetV1RaritiesInternalServerError().WithPayload(errorMessage)
			}

			mappedCardRarities := mapRaritiesToExternal(cardRarities)

			return rarities.NewGetV1RaritiesOK().WithPayload(mappedCardRarities)
		},
	)
}

func mapRaritiesToExternal(Rarities []domain.Rarity) []*models.Rarities {
	var mappedRarities []*models.Rarities
	for _, rarity := range Rarities {
		var c models.Rarities
		c.ID = int64(rarity.ID)
		c.Name = rarity.Name

		for _, cost := range rarity.CraftingCost {
			c.Craftingcost = append(c.Craftingcost, int64(cost))
		}

		for _, dustValue := range rarity.DustValue {
			c.Dustvalue = append(c.Dustvalue, int64(dustValue))
		}

		mappedRarities = append(mappedRarities, &c)
	}

	return mappedRarities
}
