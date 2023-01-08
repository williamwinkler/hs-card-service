package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/codegen/restapi"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// hsClient, err := clients.NewHsClient()
	// if err != nil {
	// 	log.Fatalf("Failed to start HsClient: %v", err)
	// }

	// cardRepository, err := repositories.NewCardRepository()
	// if err != nil {
	// 	log.Fatalf("Failed to start cardRepository: %v", err)
	// }

	// cardService := application.NewCardService(hsClient, cardRepository)

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

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Program Ended")
}
