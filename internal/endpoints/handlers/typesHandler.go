package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	types "github.com/williamwinkler/hs-card-service/codegen/restapi/operations/types"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/logging"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
	"github.com/williamwinkler/hs-card-service/internal/observability"
	"go.opentelemetry.io/otel/attribute"
)

type TypesHandler struct {
	api      *operations.HearthstoneCardServiceAPI
	typeRepo *repositories.TypeRepository
}

func NewTypesHandler(api *operations.HearthstoneCardServiceAPI, typeRepo *repositories.TypeRepository) *TypesHandler {
	return &TypesHandler{
		api:      api,
		typeRepo: typeRepo,
	}
}

func (i *TypesHandler) SetupHandler() {
	i.api.TypesGetTypesHandler = types.GetTypesHandlerFunc(
		func(req types.GetTypesParams) middleware.Responder {
			ctx, span := observability.Start(req.HTTPRequest.Context(), "metadata.get", attribute.String("app.metadata.resource", "type"))
			defer span.End()
			defer logging.Debugf(ctx, "Handled GET /types request")

			cardTypes, err := i.typeRepo.FindAll(ctx)
			if err != nil {
				observability.Fail(span, "database_failed")
				logging.Errorf(ctx, "Error occurred in GET /Types: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				return types.NewGetTypesInternalServerError().WithPayload(errorMessage)
			}

			mappedCardTypes := mapTypesToExternal(cardTypes)

			return types.NewGetTypesOK().WithPayload(mappedCardTypes)
		},
	)
}

func mapTypesToExternal(Types []domain.Type) []*models.Types {
	var mappedTypes []*models.Types
	for _, set := range Types {
		var c models.Types
		c.ID = int64(set.ID)
		c.Name = set.Name

		mappedTypes = append(mappedTypes, &c)
	}

	return mappedTypes
}
