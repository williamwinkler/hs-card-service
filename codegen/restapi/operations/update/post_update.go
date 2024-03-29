// Code generated by go-swagger; DO NOT EDIT.

package update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUpdateHandlerFunc turns a function with the right signature into a post update handler
type PostUpdateHandlerFunc func(PostUpdateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUpdateHandlerFunc) Handle(params PostUpdateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostUpdateHandler interface for that can handle valid post update params
type PostUpdateHandler interface {
	Handle(PostUpdateParams, interface{}) middleware.Responder
}

// NewPostUpdate creates a new http.Handler for the post update operation
func NewPostUpdate(ctx *middleware.Context, handler PostUpdateHandler) *PostUpdate {
	return &PostUpdate{Context: ctx, Handler: handler}
}

/*
	PostUpdate swagger:route POST /update update postUpdate

Checks for updates to Hearthstone, if there are any, they are applied to the database
*/
type PostUpdate struct {
	Context *middleware.Context
	Handler PostUpdateHandler
}

func (o *PostUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostUpdateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
