// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/williamwinkler/hs-card-service/codegen/models"
)

// GetTypesOKCode is the HTTP code returned for type GetTypesOK
const GetTypesOKCode int = 200

/*
GetTypesOK Returns all types

swagger:response getTypesOK
*/
type GetTypesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Types `json:"body,omitempty"`
}

// NewGetTypesOK creates GetTypesOK with default headers values
func NewGetTypesOK() *GetTypesOK {

	return &GetTypesOK{}
}

// WithPayload adds the payload to the get types o k response
func (o *GetTypesOK) WithPayload(payload []*models.Types) *GetTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get types o k response
func (o *GetTypesOK) SetPayload(payload []*models.Types) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Types, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTypesInternalServerErrorCode is the HTTP code returned for type GetTypesInternalServerError
const GetTypesInternalServerErrorCode int = 500

/*
GetTypesInternalServerError Something went wrong internally

swagger:response getTypesInternalServerError
*/
type GetTypesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTypesInternalServerError creates GetTypesInternalServerError with default headers values
func NewGetTypesInternalServerError() *GetTypesInternalServerError {

	return &GetTypesInternalServerError{}
}

// WithPayload adds the payload to the get types internal server error response
func (o *GetTypesInternalServerError) WithPayload(payload *models.Error) *GetTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get types internal server error response
func (o *GetTypesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
