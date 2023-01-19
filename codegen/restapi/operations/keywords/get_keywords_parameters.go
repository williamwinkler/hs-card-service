// Code generated by go-swagger; DO NOT EDIT.

package keywords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetKeywordsParams creates a new GetKeywordsParams object
//
// There are no default values defined in the spec.
func NewGetKeywordsParams() GetKeywordsParams {

	return GetKeywordsParams{}
}

// GetKeywordsParams contains all the bound params for the get keywords operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetKeywords
type GetKeywordsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetKeywordsParams() beforehand.
func (o *GetKeywordsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}