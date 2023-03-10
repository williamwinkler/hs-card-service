// Code generated by go-swagger; DO NOT EDIT.

package classes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/williamwinkler/hs-card-service/codegen/models"
)

// GetClassesOKCode is the HTTP code returned for type GetClassesOK
const GetClassesOKCode int = 200

/*
GetClassesOK Returns all classes

swagger:response getClassesOK
*/
type GetClassesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Classes `json:"body,omitempty"`
}

// NewGetClassesOK creates GetClassesOK with default headers values
func NewGetClassesOK() *GetClassesOK {

	return &GetClassesOK{}
}

// WithPayload adds the payload to the get classes o k response
func (o *GetClassesOK) WithPayload(payload []*models.Classes) *GetClassesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get classes o k response
func (o *GetClassesOK) SetPayload(payload []*models.Classes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClassesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Classes, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetClassesInternalServerErrorCode is the HTTP code returned for type GetClassesInternalServerError
const GetClassesInternalServerErrorCode int = 500

/*
GetClassesInternalServerError Something went wrong internally

swagger:response getClassesInternalServerError
*/
type GetClassesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetClassesInternalServerError creates GetClassesInternalServerError with default headers values
func NewGetClassesInternalServerError() *GetClassesInternalServerError {

	return &GetClassesInternalServerError{}
}

// WithPayload adds the payload to the get classes internal server error response
func (o *GetClassesInternalServerError) WithPayload(payload *models.Error) *GetClassesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get classes internal server error response
func (o *GetClassesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClassesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
