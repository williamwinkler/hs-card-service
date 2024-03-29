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

var portFlag = flag.Int("port", 8080, "Port to run this service on")

type RestServer struct {
	cardRepo       *repositories.CardRepository
	updateMetaRepo *repositories.UpdateMetaRepository
	setRepo        *repositories.SetRepository
	typeRepo       *repositories.TypeRepository
	classRepo      *repositories.ClassRepository
	rarityRepo     *repositories.RarityRepository
	keywordRepo    *repositories.KeywordRepository
	cardService    *application.CardService
	setService     *application.SetService
	classService   *application.ClassService
	rarityService  *application.RarityService
	typeService    *application.TypeService
	keywordService *application.KeywordService
}

func NewRestServer(
	cardRepo *repositories.CardRepository,
	updateMetaRepo *repositories.UpdateMetaRepository,
	setRepo *repositories.SetRepository,
	typeRepo *repositories.TypeRepository,
	classRepo *repositories.ClassRepository,
	rarityRepo *repositories.RarityRepository,
	keywordRepo *repositories.KeywordRepository,
	cardService *application.CardService,
	setService *application.SetService,
	classService *application.ClassService,
	rarityService *application.RarityService,
	typeService *application.TypeService,
	keywordService *application.KeywordService,

) *RestServer {
	return &RestServer{
		cardService:    cardService,
		cardRepo:       cardRepo,
		setRepo:        setRepo,
		typeRepo:       typeRepo,
		classRepo:      classRepo,
		rarityRepo:     rarityRepo,
		keywordRepo:    keywordRepo,
		updateMetaRepo: updateMetaRepo,
		setService:     setService,
		classService:   classService,
		rarityService:  rarityService,
		typeService:    typeService,
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
		handlers.NewCardUpdateHandler(api, s.cardService, s.setService, s.classService, s.rarityService, s.typeService, s.keywordService),
		handlers.NewCardHandler(api, s.cardService),
		handlers.NewRichCardHandler(api, s.cardService),
		handlers.NewSetsHandler(api, s.setRepo),
		handlers.NewTypesHandler(api, s.typeRepo),
		handlers.NewClassesHandler(api, s.classRepo),
		handlers.NewRaritiesHandler(api, s.rarityRepo),
		handlers.NewKeywordsHandler(api, s.keywordRepo),
	}
	inizializeHandlers(handlers)

	server.ConfigureAPI()

	//serve API
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
