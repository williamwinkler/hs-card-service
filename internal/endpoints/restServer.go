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
	cardService *application.CardService
	cardRepo    *repositories.CardRepository
}

func NewRestServer(
	cardService *application.CardService,
	cardRepo *repositories.CardRepository,
) *RestServer {
	return &RestServer{
		cardService: cardService,
		cardRepo:    cardRepo,
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
		handlers.NewInfoHandler(api, s.cardRepo),
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
