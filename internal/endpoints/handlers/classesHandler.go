package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/classes"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/logging"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
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
			ctx, span := observability.Start(req.HTTPRequest.Context(), "metadata.get", attribute.String("app.metadata.resource", "class"))
			defer span.End()
			defer logging.Debugf(ctx, "Handled GET /classes request")

			cardClasses, err := i.classRepo.FindAll(ctx)
			if err != nil {
				observability.Fail(span, "database_failed")
				logging.Errorf(ctx, "Error occurred in GET /classes: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				return classes.NewGetClassesInternalServerError().WithPayload(errorMessage)
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
