package handlers

import (
	"log"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/info"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

var systemStartTime = time.Now()

type InfoHandler struct {
	api            *operations.HearthstoneCardServiceAPI
	cardRepo       *repositories.CardRepository
	updateMetaRepo *repositories.UpdateMetaRepository
}

func NewInfoHandler(api *operations.HearthstoneCardServiceAPI, cardRepo *repositories.CardRepository, updateMetaRepo *repositories.UpdateMetaRepository) *InfoHandler {
	return &InfoHandler{
		api:            api,
		cardRepo:       cardRepo,
		updateMetaRepo: updateMetaRepo,
	}
}

func (i *InfoHandler) SetupHandler() {
	i.api.InfoGetInfoHandler = info.GetInfoHandlerFunc(
		func(gp info.GetInfoParams) middleware.Responder {
			defer log.Printf("Handled %s request", gp.HTTPRequest.URL)

			count, err := i.cardRepo.Count()
			if err != nil {
				log.Printf("Error occurred in GET /info: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				return info.NewGetInfoInternalServerError().WithPayload(errorMessage)
			}

			newestUpdate, err := i.updateMetaRepo.FindNewest()
			if err != nil {
				log.Printf("Error occurred in GET /info: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				return info.NewGetInfoInternalServerError().WithPayload(errorMessage)
			}

			infoResponse := models.Info{
				AmountOfCards:   count,
				LastUpdate:      strfmt.DateTime(newestUpdate.Updated),
				SystemStartTime: strfmt.DateTime(systemStartTime),
			}
			return info.NewGetInfoOK().WithPayload(&infoResponse)
		},
	)
}
