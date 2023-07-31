package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/endpoints"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/clients"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Found no .env file")
	}

	database, err := migrations.SetupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	hsClient, err := clients.NewHsClient()
	if err != nil {
		log.Fatalf("Failed to start HsClient: %v", err)
	}

	// setup repositories
	cardRepository := repositories.NewCardRepository(database.Db)
	cardMetaRepository := repositories.NewUpdateMetaRepository(database.Db)
	setRepository := repositories.NewSetRepository(database.Db)
	classRepository := repositories.NewClassRepository(database.Db)
	rarityRepository := repositories.NewRarityRepository(database.Db)
	typeRepository := repositories.NewTypeRepository(database.Db)
	keywordRepository := repositories.NewKeywordRepository(database.Db)

	// setup services
	cardService := application.NewCardService(hsClient, cardRepository, cardMetaRepository)
	setService := application.NewSetService(setRepository, hsClient)
	classService := application.NewClassService(classRepository, hsClient)
	rarityService := application.NewRarityService(rarityRepository, hsClient)
	typeService := application.NewTypeService(typeRepository, hsClient)
	keywordService := application.NewKeywordService(keywordRepository, hsClient)

	// setup server
	restServer := endpoints.NewRestServer(
		cardRepository,
		cardMetaRepository,
		setRepository,
		typeRepository,
		classRepository,
		rarityRepository,
		keywordRepository,
		cardService,
		setService,
		classService,
		rarityService,
		typeService,
		keywordService,
	)
	restServer.StartServer()

	log.Println("Program Ended")
}
