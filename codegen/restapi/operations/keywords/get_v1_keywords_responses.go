// Code generated by go-swagger; DO NOT EDIT.

package keywords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/williamwinkler/hs-card-service/codegen/models"
)

// GetV1KeywordsOKCode is the HTTP code returned for type GetV1KeywordsOK
const GetV1KeywordsOKCode int = 200

/*
GetV1KeywordsOK Returns all keywords

swagger:response getV1KeywordsOK
*/
type GetV1KeywordsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Keywords `json:"body,omitempty"`
}

// NewGetV1KeywordsOK creates GetV1KeywordsOK with default headers values
func NewGetV1KeywordsOK() *GetV1KeywordsOK {

	return &GetV1KeywordsOK{}
}

// WithPayload adds the payload to the get v1 keywords o k response
func (o *GetV1KeywordsOK) WithPayload(payload []*models.Keywords) *GetV1KeywordsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 keywords o k response
func (o *GetV1KeywordsOK) SetPayload(payload []*models.Keywords) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1KeywordsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Keywords, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1KeywordsInternalServerErrorCode is the HTTP code returned for type GetV1KeywordsInternalServerError
const GetV1KeywordsInternalServerErrorCode int = 500

/*
GetV1KeywordsInternalServerError Something went wrong internally

swagger:response getV1KeywordsInternalServerError
*/
type GetV1KeywordsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetV1KeywordsInternalServerError creates GetV1KeywordsInternalServerError with default headers values
func NewGetV1KeywordsInternalServerError() *GetV1KeywordsInternalServerError {

	return &GetV1KeywordsInternalServerError{}
}

// WithPayload adds the payload to the get v1 keywords internal server error response
func (o *GetV1KeywordsInternalServerError) WithPayload(payload *models.Error) *GetV1KeywordsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 keywords internal server error response
func (o *GetV1KeywordsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1KeywordsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
