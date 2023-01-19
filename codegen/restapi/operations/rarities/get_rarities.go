// Code generated by go-swagger; DO NOT EDIT.

package rarities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRaritiesHandlerFunc turns a function with the right signature into a get rarities handler
type GetRaritiesHandlerFunc func(GetRaritiesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRaritiesHandlerFunc) Handle(params GetRaritiesParams) middleware.Responder {
	return fn(params)
}

// GetRaritiesHandler interface for that can handle valid get rarities params
type GetRaritiesHandler interface {
	Handle(GetRaritiesParams) middleware.Responder
}

// NewGetRarities creates a new http.Handler for the get rarities operation
func NewGetRarities(ctx *middleware.Context, handler GetRaritiesHandler) *GetRarities {
	return &GetRarities{Context: ctx, Handler: handler}
}

/*
	GetRarities swagger:route GET /rarities rarities getRarities

Serves the different rarities a card can have. Fx "Common" or "Legendary"
*/
type GetRarities struct {
	Context *middleware.Context
	Handler GetRaritiesHandler
}

func (o *GetRarities) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRaritiesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}