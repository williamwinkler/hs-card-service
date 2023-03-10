// Code generated by go-swagger; DO NOT EDIT.

package cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRichcardsHandlerFunc turns a function with the right signature into a get richcards handler
type GetRichcardsHandlerFunc func(GetRichcardsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRichcardsHandlerFunc) Handle(params GetRichcardsParams) middleware.Responder {
	return fn(params)
}

// GetRichcardsHandler interface for that can handle valid get richcards params
type GetRichcardsHandler interface {
	Handle(GetRichcardsParams) middleware.Responder
}

// NewGetRichcards creates a new http.Handler for the get richcards operation
func NewGetRichcards(ctx *middleware.Context, handler GetRichcardsHandler) *GetRichcards {
	return &GetRichcards{Context: ctx, Handler: handler}
}

/*
	GetRichcards swagger:route GET /richcards cards getRichcards

Rich cards contains names instead of ids of fx CardType "Minion", "Spell", "Secret" etc
*/
type GetRichcards struct {
	Context *middleware.Context
	Handler GetRichcardsHandler
}

func (o *GetRichcards) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRichcardsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
