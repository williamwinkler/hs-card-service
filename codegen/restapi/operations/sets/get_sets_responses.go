// Code generated by go-swagger; DO NOT EDIT.

package sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/williamwinkler/hs-card-service/codegen/models"
)

// GetSetsOKCode is the HTTP code returned for type GetSetsOK
const GetSetsOKCode int = 200

/*
GetSetsOK Returns all sets

swagger:response getSetsOK
*/
type GetSetsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Sets `json:"body,omitempty"`
}

// NewGetSetsOK creates GetSetsOK with default headers values
func NewGetSetsOK() *GetSetsOK {

	return &GetSetsOK{}
}

// WithPayload adds the payload to the get sets o k response
func (o *GetSetsOK) WithPayload(payload []*models.Sets) *GetSetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sets o k response
func (o *GetSetsOK) SetPayload(payload []*models.Sets) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Sets, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSetsInternalServerErrorCode is the HTTP code returned for type GetSetsInternalServerError
const GetSetsInternalServerErrorCode int = 500

/*
GetSetsInternalServerError Something went wrong internally

swagger:response getSetsInternalServerError
*/
type GetSetsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSetsInternalServerError creates GetSetsInternalServerError with default headers values
func NewGetSetsInternalServerError() *GetSetsInternalServerError {

	return &GetSetsInternalServerError{}
}

// WithPayload adds the payload to the get sets internal server error response
func (o *GetSetsInternalServerError) WithPayload(payload *models.Error) *GetSetsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sets internal server error response
func (o *GetSetsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSetsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}