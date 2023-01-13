// Code generated by go-swagger; DO NOT EDIT.

package update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/williamwinkler/hs-card-service/codegen/models"
)

// PostUpdateOKCode is the HTTP code returned for type PostUpdateOK
const PostUpdateOKCode int = 200

/*
PostUpdateOK OK

swagger:response postUpdateOK
*/
type PostUpdateOK struct {
}

// NewPostUpdateOK creates PostUpdateOK with default headers values
func NewPostUpdateOK() *PostUpdateOK {

	return &PostUpdateOK{}
}

// WriteResponse to the client
func (o *PostUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostUpdateUnauthorizedCode is the HTTP code returned for type PostUpdateUnauthorized
const PostUpdateUnauthorizedCode int = 401

/*
PostUpdateUnauthorized Unauthorized request

swagger:response postUpdateUnauthorized
*/
type PostUpdateUnauthorized struct {
}

// NewPostUpdateUnauthorized creates PostUpdateUnauthorized with default headers values
func NewPostUpdateUnauthorized() *PostUpdateUnauthorized {

	return &PostUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *PostUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostUpdateInternalServerErrorCode is the HTTP code returned for type PostUpdateInternalServerError
const PostUpdateInternalServerErrorCode int = 500

/*
PostUpdateInternalServerError Something went wrong internally

swagger:response postUpdateInternalServerError
*/
type PostUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUpdateInternalServerError creates PostUpdateInternalServerError with default headers values
func NewPostUpdateInternalServerError() *PostUpdateInternalServerError {

	return &PostUpdateInternalServerError{}
}

// WithPayload adds the payload to the post update internal server error response
func (o *PostUpdateInternalServerError) WithPayload(payload *models.Error) *PostUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post update internal server error response
func (o *PostUpdateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}