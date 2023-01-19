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

func NewSetHandler(api *operations.HearthstoneCardServiceAPI, setRepo *repositories.SetRepository) *SetsHandler {
	return &SetsHandler{
		api:     api,
		setRepo: setRepo,
	}
}

func (i *SetsHandler) SetupHandler() {
	i.api.SetsGetSetsHandler = sets.GetSetsHandlerFunc(
		func(req sets.GetSetsParams) middleware.Responder {
			defer log.Printf("Handled %s request", req.HTTPRequest.URL)

			cardSets, err := i.setRepo.FindAll()
			if err != nil {
				log.Printf("Error occurred in GET /sets: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				sets.NewGetSetsInternalServerError().WithPayload(errorMessage)
			}

			mappedCardSets := mapSetsToExternal(cardSets)

			return sets.NewGetSetsOK().WithPayload(mappedCardSets)
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
