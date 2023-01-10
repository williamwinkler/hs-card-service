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
	api      *operations.HearthstoneCardServiceAPI
	cardRepo *repositories.CardRepository
}

func NewInfoHandler(api *operations.HearthstoneCardServiceAPI, cardRepo *repositories.CardRepository) *InfoHandler {
	return &InfoHandler{
		api:      api,
		cardRepo: cardRepo,
	}
}

func (i *InfoHandler) SetupHandler() {
	i.api.InfoGetHandler = info.GetHandlerFunc(
		func(gp info.GetParams) middleware.Responder {
			defer log.Println("Handled GetInfo request")

			count, err := i.cardRepo.Count()
			if err != nil {
				log.Printf("Error occurred in GET /info: %v", err)
				errorMessage := utils.CreateErrorMessage(500, "Something went wrong getting with getting card count")
				return info.NewGetInternalServerError().WithPayload(errorMessage)
			}

			infoResponse := models.Info{
				AmountOfCards:   count,
				LastUpdate:      strfmt.DateTime(time.Now()),
				SystemStartTime: strfmt.DateTime(systemStartTime),
			}
			return info.NewGetOK().WithPayload(&infoResponse)
		},
	)

}
