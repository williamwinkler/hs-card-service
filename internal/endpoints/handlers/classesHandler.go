package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/classes"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

type ClassesHandler struct {
	api       *operations.HearthstoneCardServiceAPI
	classRepo *repositories.ClassRepository
}

func NewClassesHandler(api *operations.HearthstoneCardServiceAPI, classRepo *repositories.ClassRepository) *ClassesHandler {
	return &ClassesHandler{
		api:       api,
		classRepo: classRepo,
	}
}

func (i *ClassesHandler) SetupHandler() {
	i.api.ClassesGetClassesHandler = classes.GetClassesHandlerFunc(
		func(req classes.GetClassesParams) middleware.Responder {
			defer log.Printf("Handled %s request", req.HTTPRequest.URL)

			cardClasses, err := i.classRepo.FindAll()
			if err != nil {
				log.Printf("Error occurred in GET /classes: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				classes.NewGetClassesInternalServerError().WithPayload(errorMessage)
			}

			mappedCardClasses := mapClassesToExternal(cardClasses)

			return classes.NewGetClassesOK().WithPayload(mappedCardClasses)
		},
	)
}

func mapClassesToExternal(Classes []domain.Class) []*models.Classes {
	var mappedClasses []*models.Classes
	for _, set := range Classes {
		var c models.Classes
		c.ID = int64(set.ID)
		c.Name = set.Name

		mappedClasses = append(mappedClasses, &c)
	}

	return mappedClasses
}
