// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/classes"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/info"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/keywords"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/rarities"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/sets"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/types"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/update"
)

//go:generate swagger generate server --target ../../codegen --name HearthstoneCardService --spec ../../api/swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.HearthstoneCardServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HearthstoneCardServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CardsGetCardsHandler == nil {
		api.CardsGetCardsHandler = cards.GetCardsHandlerFunc(func(params cards.GetCardsParams) middleware.Responder {
			return middleware.NotImplemented("operation cards.GetCards has not yet been implemented")
		})
	}
	if api.ClassesGetClassesHandler == nil {
		api.ClassesGetClassesHandler = classes.GetClassesHandlerFunc(func(params classes.GetClassesParams) middleware.Responder {
			return middleware.NotImplemented("operation classes.GetClasses has not yet been implemented")
		})
	}
	if api.InfoGetInfoHandler == nil {
		api.InfoGetInfoHandler = info.GetInfoHandlerFunc(func(params info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation info.GetInfo has not yet been implemented")
		})
	}
	if api.KeywordsGetKeywordsHandler == nil {
		api.KeywordsGetKeywordsHandler = keywords.GetKeywordsHandlerFunc(func(params keywords.GetKeywordsParams) middleware.Responder {
			return middleware.NotImplemented("operation keywords.GetKeywords has not yet been implemented")
		})
	}
	if api.RaritiesGetRaritiesHandler == nil {
		api.RaritiesGetRaritiesHandler = rarities.GetRaritiesHandlerFunc(func(params rarities.GetRaritiesParams) middleware.Responder {
			return middleware.NotImplemented("operation rarities.GetRarities has not yet been implemented")
		})
	}
	if api.CardsGetRichcardsHandler == nil {
		api.CardsGetRichcardsHandler = cards.GetRichcardsHandlerFunc(func(params cards.GetRichcardsParams) middleware.Responder {
			return middleware.NotImplemented("operation cards.GetRichcards has not yet been implemented")
		})
	}
	if api.SetsGetSetsHandler == nil {
		api.SetsGetSetsHandler = sets.GetSetsHandlerFunc(func(params sets.GetSetsParams) middleware.Responder {
			return middleware.NotImplemented("operation sets.GetSets has not yet been implemented")
		})
	}
	if api.TypesGetTypesHandler == nil {
		api.TypesGetTypesHandler = types.GetTypesHandlerFunc(func(params types.GetTypesParams) middleware.Responder {
			return middleware.NotImplemented("operation types.GetTypes has not yet been implemented")
		})
	}
	if api.UpdatePostUpdateHandler == nil {
		api.UpdatePostUpdateHandler = update.PostUpdateHandlerFunc(func(params update.PostUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation update.PostUpdate has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{AllowedOrigins: []string{"http://localhost:3000"}})
	return c.Handler(handler)
}
