// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1TypesHandlerFunc turns a function with the right signature into a get v1 types handler
type GetV1TypesHandlerFunc func(GetV1TypesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1TypesHandlerFunc) Handle(params GetV1TypesParams) middleware.Responder {
	return fn(params)
}

// GetV1TypesHandler interface for that can handle valid get v1 types params
type GetV1TypesHandler interface {
	Handle(GetV1TypesParams) middleware.Responder
}

// NewGetV1Types creates a new http.Handler for the get v1 types operation
func NewGetV1Types(ctx *middleware.Context, handler GetV1TypesHandler) *GetV1Types {
	return &GetV1Types{Context: ctx, Handler: handler}
}

/*
	GetV1Types swagger:route GET /v1/types types getV1Types

Serves the different types cards can be. Fx "Minion" or "Spell"
*/
type GetV1Types struct {
	Context *middleware.Context
	Handler GetV1TypesHandler
}

func (o *GetV1Types) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetV1TypesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}