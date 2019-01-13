// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gitpods/gitpods/pkg/api/v1/models"
)

// GetUserMeOKCode is the HTTP code returned for type GetUserMeOK
const GetUserMeOKCode int = 200

/*GetUserMeOK The current authenticated user

swagger:response getUserMeOK
*/
type GetUserMeOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserMeOK creates GetUserMeOK with default headers values
func NewGetUserMeOK() *GetUserMeOK {

	return &GetUserMeOK{}
}

// WithPayload adds the payload to the get user me o k response
func (o *GetUserMeOK) WithPayload(payload *models.User) *GetUserMeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user me o k response
func (o *GetUserMeOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserMeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUserMeDefault unexpected error

swagger:response getUserMeDefault
*/
type GetUserMeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserMeDefault creates GetUserMeDefault with default headers values
func NewGetUserMeDefault(code int) *GetUserMeDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserMeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user me default response
func (o *GetUserMeDefault) WithStatusCode(code int) *GetUserMeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user me default response
func (o *GetUserMeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user me default response
func (o *GetUserMeDefault) WithPayload(payload *models.Error) *GetUserMeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user me default response
func (o *GetUserMeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserMeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
