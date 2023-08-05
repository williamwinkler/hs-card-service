package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/sets"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

type SetsHandler struct {
	api     *operations.HearthstoneCardServiceAPI
	setRepo *repositories.SetRepository
}

func NewSetsHandler(api *operations.HearthstoneCardServiceAPI, setRepo *repositories.SetRepository) *SetsHandler {
	return &SetsHandler{
		api:     api,
		setRepo: setRepo,
	}
}

func (i *SetsHandler) SetupHandler() {
	i.api.SetsGetV1SetsHandler = sets.GetV1SetsHandlerFunc(
		func(req sets.GetV1SetsParams) middleware.Responder {
			defer log.Printf("Handled %s request", req.HTTPRequest.URL)

			cardSets, err := i.setRepo.FindAll()
			if err != nil {
				log.Printf("Error occurred in GET /sets: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				sets.NewGetV1SetsInternalServerError().WithPayload(errorMessage)
			}

			mappedCardSets := mapSetsToExternal(cardSets)

			return sets.NewGetV1SetsOK().WithPayload(mappedCardSets)
		},
	)
}

func mapSetsToExternal(sets []domain.Set) []*models.Sets {
	var mappedSets []*models.Sets
	for _, set := range sets {
		var c models.Sets
		c.ID = int64(set.ID)
		c.Name = set.Name
		c.Type = set.Type

		mappedSets = append(mappedSets, &c)
	}

	return mappedSets
}
