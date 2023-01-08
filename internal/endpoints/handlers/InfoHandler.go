package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/info"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

type InfoHandler struct {
	api      *operations.HearthstoneCardServiceAPI
	cardRepo *repositories.CardRepository
}

func NewInfoHandler(api *operations.HearthstoneCardServiceAPI, cardRepo *repositories.CardRepository) *InfoHandler {
	return &InfoHandler{
		api:      api,
		cardRepo: cardRepo,
	}
}

func (i *InfoHandler) SetupInfoHandler() {
	i.api.InfoGetHandler = info.GetHandlerFunc(
		func(gp info.GetParams) middleware.Responder {

			count, err := i.cardRepo.Count()
			if err != nil {
				errorMessage := utils.CreateErrorMessage(500, "Something went wrong")
				return info.NewGetInternalServerError().WithPayload(errorMessage)
			}

			infoResponse := models.Info{
				AmountOfCards:       count,
				TimeSinceLastUpdate: 2,
				Uptime:              56,
			}
			return info.NewGetOK().WithPayload(&infoResponse)
		},
	)

}
