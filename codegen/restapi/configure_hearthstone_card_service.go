// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/classes"
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

	// Add basic authentication to the API
	api.BasicAuthAuth = func(username, password string) (interface{}, error) {
		validUsername := os.Getenv("BASIC_AUTH_USERNAME")
		validPassword := os.Getenv("BASIC_AUTH_PASSWORD")
		if username == validUsername && password == validPassword {
			return username, nil
		}
		return nil, errors.New(401, "Unauthorized")
	}

	if api.CardsGetV1CardsHandler == nil {
		api.CardsGetV1CardsHandler = cards.GetV1CardsHandlerFunc(func(params cards.GetV1CardsParams) middleware.Responder {
			return middleware.NotImplemented("operation cards.GetV1Cards has not yet been implemented")
		})
	}
	if api.ClassesGetV1ClassesHandler == nil {
		api.ClassesGetV1ClassesHandler = classes.GetV1ClassesHandlerFunc(func(params classes.GetV1ClassesParams) middleware.Responder {
			return middleware.NotImplemented("operation classes.GetV1Classes has not yet been implemented")
		})
	}

	if api.KeywordsGetV1KeywordsHandler == nil {
		api.KeywordsGetV1KeywordsHandler = keywords.GetV1KeywordsHandlerFunc(func(params keywords.GetV1KeywordsParams) middleware.Responder {
			return middleware.NotImplemented("operation keywords.GetV1Keywords has not yet been implemented")
		})
	}
	if api.RaritiesGetV1RaritiesHandler == nil {
		api.RaritiesGetV1RaritiesHandler = rarities.GetV1RaritiesHandlerFunc(func(params rarities.GetV1RaritiesParams) middleware.Responder {
			return middleware.NotImplemented("operation rarities.GetV1Rarities has not yet been implemented")
		})
	}
	if api.CardsGetV1RichcardsHandler == nil {
		api.CardsGetV1RichcardsHandler = cards.GetV1RichcardsHandlerFunc(func(params cards.GetV1RichcardsParams) middleware.Responder {
			return middleware.NotImplemented("operation cards.GetV1Richcards has not yet been implemented")
		})
	}
	if api.SetsGetV1SetsHandler == nil {
		api.SetsGetV1SetsHandler = sets.GetV1SetsHandlerFunc(func(params sets.GetV1SetsParams) middleware.Responder {
			return middleware.NotImplemented("operation sets.GetV1Sets has not yet been implemented")
		})
	}
	if api.TypesGetV1TypesHandler == nil {
		api.TypesGetV1TypesHandler = types.GetV1TypesHandlerFunc(func(params types.GetV1TypesParams) middleware.Responder {
			return middleware.NotImplemented("operation types.GetV1Types has not yet been implemented")
		})
	}
	if api.UpdatePostUpdateHandler == nil {
		api.UpdatePostUpdateHandler = update.PostUpdateHandlerFunc(func(params update.PostUpdateParams, i interface{}) middleware.Responder {
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

	lmt := tollbooth.NewLimiter(15, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})
	lmt.SetMessage("You have reached maximum request limit.")

	return tollbooth.LimitFuncHandler(lmt, handler.ServeHTTP)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{AllowedOrigins: []string{"http://localhost:3000"}})
	return c.Handler(handler)
}
