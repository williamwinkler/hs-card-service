package endpoints

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/williamwinkler/hs-card-service/codegen/restapi"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/endpoints/handlers"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

type RestServer struct {
	cardRepo       *repositories.CardRepository
	updateMetaRepo *repositories.UpdateMetaRepository
	cardService    *application.CardService
	setService     *application.SetService
	classService   *application.ClassService
	rarityService  *application.RarityService
	keywordService *application.KeywordService
}

func NewRestServer(
	cardRepo *repositories.CardRepository,
	updateMetaRepo *repositories.UpdateMetaRepository,
	cardService *application.CardService,
	setService *application.SetService,
	classService *application.ClassService,
	rarityService *application.RarityService,
	keywordService *application.KeywordService,

) *RestServer {
	return &RestServer{
		cardService:    cardService,
		cardRepo:       cardRepo,
		updateMetaRepo: updateMetaRepo,
		setService:     setService,
		classService:   classService,
		rarityService:  rarityService,
		keywordService: keywordService,
	}
}

func (s *RestServer) StartServer() {
	// load swagger spec
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewHearthstoneCardServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	// server.ConfigureAPI()

	// parse flags
	flag.Parse()
	server.Port = *portFlag

	// inizalize handlers
	handlers := []Handler{
		handlers.NewInfoHandler(api, s.cardRepo, s.updateMetaRepo),
		handlers.NewCardUpdateHandler(api, s.cardService, s.setService, s.classService, s.rarityService, s.keywordService),
		handlers.NewCardHandler(api, s.cardService),
	}
	inizializeHandlers(handlers)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

type Handler interface {
	SetupHandler()
}

func inizializeHandlers(handlers []Handler) {
	for _, handler := range handlers {
		handler.SetupHandler()
	}
}
