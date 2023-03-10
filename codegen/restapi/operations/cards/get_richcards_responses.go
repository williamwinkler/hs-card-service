// Code generated by go-swagger; DO NOT EDIT.

package cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/williamwinkler/hs-card-service/codegen/models"
)

// GetRichcardsOKCode is the HTTP code returned for type GetRichcardsOK
const GetRichcardsOKCode int = 200

/*
GetRichcardsOK Returns the cards based on query. If there is no query, cards will be returned based on their manaCost in ascending order.

swagger:response getRichcardsOK
*/
type GetRichcardsOK struct {

	/*
	  In: Body
	*/
	Payload *models.RichCards `json:"body,omitempty"`
}

// NewGetRichcardsOK creates GetRichcardsOK with default headers values
func NewGetRichcardsOK() *GetRichcardsOK {

	return &GetRichcardsOK{}
}

// WithPayload adds the payload to the get richcards o k response
func (o *GetRichcardsOK) WithPayload(payload *models.RichCards) *GetRichcardsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get richcards o k response
func (o *GetRichcardsOK) SetPayload(payload *models.RichCards) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRichcardsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRichcardsInternalServerErrorCode is the HTTP code returned for type GetRichcardsInternalServerError
const GetRichcardsInternalServerErrorCode int = 500

/*
GetRichcardsInternalServerError Something went wrong internally

swagger:response getRichcardsInternalServerError
*/
type GetRichcardsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRichcardsInternalServerError creates GetRichcardsInternalServerError with default headers values
func NewGetRichcardsInternalServerError() *GetRichcardsInternalServerError {

	return &GetRichcardsInternalServerError{}
}

// WithPayload adds the payload to the get richcards internal server error response
func (o *GetRichcardsInternalServerError) WithPayload(payload *models.Error) *GetRichcardsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get richcards internal server error response
func (o *GetRichcardsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRichcardsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
