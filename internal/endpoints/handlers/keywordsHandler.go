package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/williamwinkler/hs-card-service/codegen/models"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/keywords"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers/utils"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

type KeywordsHandler struct {
	api         *operations.HearthstoneCardServiceAPI
	keywordRepo *repositories.KeywordRepository
}

func NewKeywordsHandler(api *operations.HearthstoneCardServiceAPI, keywordRepo *repositories.KeywordRepository) *KeywordsHandler {
	return &KeywordsHandler{
		api:         api,
		keywordRepo: keywordRepo,
	}
}

func (i *KeywordsHandler) SetupHandler() {
	i.api.KeywordsGetKeywordsHandler = keywords.GetKeywordsHandlerFunc(
		func(req keywords.GetKeywordsParams) middleware.Responder {
			defer log.Printf("Handled %s request", req.HTTPRequest.URL)

			cardKeywords, err := i.keywordRepo.FindAll()
			if err != nil {
				log.Printf("Error occurred in GET /keywords: %v", err)
				errorMessage := utils.CreateErrorMessage(500)
				keywords.NewGetKeywordsInternalServerError().WithPayload(errorMessage)
			}

			mappedCardKeywords := mapKeywordsToExternal(cardKeywords)

			return keywords.NewGetKeywordsOK().WithPayload(mappedCardKeywords)
		},
	)
}

func mapKeywordsToExternal(Keywords []domain.Keyword) []*models.Keywords {
	var mappedKeywords []*models.Keywords
	for _, keyword := range Keywords {
		var c models.Keywords
		c.ID = int64(keyword.ID)
		c.Name = keyword.Name
		c.Text = keyword.Text

		mappedKeywords = append(mappedKeywords, &c)
	}

	return mappedKeywords
}
