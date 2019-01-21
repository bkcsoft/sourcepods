// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sourcepods/sourcepods/pkg/api/v1/models"
)

// ListUsersOKCode is the HTTP code returned for type ListUsersOK
const ListUsersOKCode int = 200

/*ListUsersOK An array of all users

swagger:response listUsersOK
*/
type ListUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewListUsersOK creates ListUsersOK with default headers values
func NewListUsersOK() *ListUsersOK {

	return &ListUsersOK{}
}

// WithPayload adds the payload to the list users o k response
func (o *ListUsersOK) WithPayload(payload []*models.User) *ListUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users o k response
func (o *ListUsersOK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ListUsersDefault unexpected error

swagger:response listUsersDefault
*/
type ListUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListUsersDefault creates ListUsersDefault with default headers values
func NewListUsersDefault(code int) *ListUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &ListUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list users default response
func (o *ListUsersDefault) WithStatusCode(code int) *ListUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list users default response
func (o *ListUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list users default response
func (o *ListUsersDefault) WithPayload(payload *models.Error) *ListUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users default response
func (o *ListUsersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
