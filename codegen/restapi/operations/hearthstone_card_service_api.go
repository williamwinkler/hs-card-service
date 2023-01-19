// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/cards"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/classes"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/info"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/keywords"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/rarities"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/sets"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/types"
	"github.com/williamwinkler/hs-card-service/codegen/restapi/operations/update"
)

// NewHearthstoneCardServiceAPI creates a new HearthstoneCardService instance
func NewHearthstoneCardServiceAPI(spec *loads.Document) *HearthstoneCardServiceAPI {
	return &HearthstoneCardServiceAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CardsGetCardsHandler: cards.GetCardsHandlerFunc(func(params cards.GetCardsParams) middleware.Responder {
			return middleware.NotImplemented("operation cards.GetCards has not yet been implemented")
		}),
		ClassesGetClassesHandler: classes.GetClassesHandlerFunc(func(params classes.GetClassesParams) middleware.Responder {
			return middleware.NotImplemented("operation classes.GetClasses has not yet been implemented")
		}),
		InfoGetInfoHandler: info.GetInfoHandlerFunc(func(params info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation info.GetInfo has not yet been implemented")
		}),
		KeywordsGetKeywordsHandler: keywords.GetKeywordsHandlerFunc(func(params keywords.GetKeywordsParams) middleware.Responder {
			return middleware.NotImplemented("operation keywords.GetKeywords has not yet been implemented")
		}),
		RaritiesGetRaritiesHandler: rarities.GetRaritiesHandlerFunc(func(params rarities.GetRaritiesParams) middleware.Responder {
			return middleware.NotImplemented("operation rarities.GetRarities has not yet been implemented")
		}),
		CardsGetRichcardsHandler: cards.GetRichcardsHandlerFunc(func(params cards.GetRichcardsParams) middleware.Responder {
			return middleware.NotImplemented("operation cards.GetRichcards has not yet been implemented")
		}),
		SetsGetSetsHandler: sets.GetSetsHandlerFunc(func(params sets.GetSetsParams) middleware.Responder {
			return middleware.NotImplemented("operation sets.GetSets has not yet been implemented")
		}),
		TypesGetTypesHandler: types.GetTypesHandlerFunc(func(params types.GetTypesParams) middleware.Responder {
			return middleware.NotImplemented("operation types.GetTypes has not yet been implemented")
		}),
		UpdatePostUpdateHandler: update.PostUpdateHandlerFunc(func(params update.PostUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation update.PostUpdate has not yet been implemented")
		}),
	}
}

/*HearthstoneCardServiceAPI Serves Hearthstone cards */
type HearthstoneCardServiceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// CardsGetCardsHandler sets the operation handler for the get cards operation
	CardsGetCardsHandler cards.GetCardsHandler
	// ClassesGetClassesHandler sets the operation handler for the get classes operation
	ClassesGetClassesHandler classes.GetClassesHandler
	// InfoGetInfoHandler sets the operation handler for the get info operation
	InfoGetInfoHandler info.GetInfoHandler
	// KeywordsGetKeywordsHandler sets the operation handler for the get keywords operation
	KeywordsGetKeywordsHandler keywords.GetKeywordsHandler
	// RaritiesGetRaritiesHandler sets the operation handler for the get rarities operation
	RaritiesGetRaritiesHandler rarities.GetRaritiesHandler
	// CardsGetRichcardsHandler sets the operation handler for the get richcards operation
	CardsGetRichcardsHandler cards.GetRichcardsHandler
	// SetsGetSetsHandler sets the operation handler for the get sets operation
	SetsGetSetsHandler sets.GetSetsHandler
	// TypesGetTypesHandler sets the operation handler for the get types operation
	TypesGetTypesHandler types.GetTypesHandler
	// UpdatePostUpdateHandler sets the operation handler for the post update operation
	UpdatePostUpdateHandler update.PostUpdateHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *HearthstoneCardServiceAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *HearthstoneCardServiceAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *HearthstoneCardServiceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *HearthstoneCardServiceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *HearthstoneCardServiceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *HearthstoneCardServiceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *HearthstoneCardServiceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *HearthstoneCardServiceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *HearthstoneCardServiceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the HearthstoneCardServiceAPI
func (o *HearthstoneCardServiceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CardsGetCardsHandler == nil {
		unregistered = append(unregistered, "cards.GetCardsHandler")
	}
	if o.ClassesGetClassesHandler == nil {
		unregistered = append(unregistered, "classes.GetClassesHandler")
	}
	if o.InfoGetInfoHandler == nil {
		unregistered = append(unregistered, "info.GetInfoHandler")
	}
	if o.KeywordsGetKeywordsHandler == nil {
		unregistered = append(unregistered, "keywords.GetKeywordsHandler")
	}
	if o.RaritiesGetRaritiesHandler == nil {
		unregistered = append(unregistered, "rarities.GetRaritiesHandler")
	}
	if o.CardsGetRichcardsHandler == nil {
		unregistered = append(unregistered, "cards.GetRichcardsHandler")
	}
	if o.SetsGetSetsHandler == nil {
		unregistered = append(unregistered, "sets.GetSetsHandler")
	}
	if o.TypesGetTypesHandler == nil {
		unregistered = append(unregistered, "types.GetTypesHandler")
	}
	if o.UpdatePostUpdateHandler == nil {
		unregistered = append(unregistered, "update.PostUpdateHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *HearthstoneCardServiceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *HearthstoneCardServiceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *HearthstoneCardServiceAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *HearthstoneCardServiceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *HearthstoneCardServiceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *HearthstoneCardServiceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the hearthstone card service API
func (o *HearthstoneCardServiceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *HearthstoneCardServiceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cards"] = cards.NewGetCards(o.context, o.CardsGetCardsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/classes"] = classes.NewGetClasses(o.context, o.ClassesGetClassesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/info"] = info.NewGetInfo(o.context, o.InfoGetInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/keywords"] = keywords.NewGetKeywords(o.context, o.KeywordsGetKeywordsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rarities"] = rarities.NewGetRarities(o.context, o.RaritiesGetRaritiesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/richcards"] = cards.NewGetRichcards(o.context, o.CardsGetRichcardsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sets"] = sets.NewGetSets(o.context, o.SetsGetSetsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/types"] = types.NewGetTypes(o.context, o.TypesGetTypesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/update"] = update.NewPostUpdate(o.context, o.UpdatePostUpdateHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *HearthstoneCardServiceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *HearthstoneCardServiceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *HearthstoneCardServiceAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *HearthstoneCardServiceAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *HearthstoneCardServiceAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
