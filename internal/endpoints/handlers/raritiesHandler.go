package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/rarities"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/logging"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
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
	i.api.RaritiesGetRaritiesHandler = rarities.GetRaritiesHandlerFunc(
		func(req rarities.GetRaritiesParams) middleware.Responder {
			ctx, span := observability.Start(req.HTTPRequest.Context(), "metadata.get", attribute.String("app.metadata.resource", "rarity"))
			defer span.End()
			defer logging.Debugf(ctx, "Handled GET /rarities request")

			cardRarities, err := i.rarityRepo.FindAll(ctx)
			if err != nil {
				observability.Fail(span, "database_failed")
				logging.Errorf(ctx, "Error occurred in GET /Rarities: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				return rarities.NewGetRaritiesInternalServerError().WithPayload(errorMessage)
			}

			mappedCardRarities := mapRaritiesToExternal(cardRarities)

			return rarities.NewGetRaritiesOK().WithPayload(mappedCardRarities)
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
