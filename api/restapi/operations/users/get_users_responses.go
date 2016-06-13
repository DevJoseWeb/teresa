package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/paas/api/models"
)

/*GetUsersOK Get users

swagger:response getUsersOK
*/
type GetUsersOK struct {

	// In: body
	Payload GetUsersOKBodyBody `json:"body,omitempty"`
}

// NewGetUsersOK creates GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

// WithPayload adds the payload to the get users o k response
func (o *GetUsersOK) WithPayload(payload GetUsersOKBodyBody) *GetUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users o k response
func (o *GetUsersOK) SetPayload(payload GetUsersOKBodyBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetUsersBadRequest User sent a bad request

swagger:response getUsersBadRequest
*/
type GetUsersBadRequest struct {

	// In: body
	Payload *models.BadRequest `json:"body,omitempty"`
}

// NewGetUsersBadRequest creates GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

// WithPayload adds the payload to the get users bad request response
func (o *GetUsersBadRequest) WithPayload(payload *models.BadRequest) *GetUsersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users bad request response
func (o *GetUsersBadRequest) SetPayload(payload *models.BadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUsersUnauthorized User not authorized

swagger:response getUsersUnauthorized
*/
type GetUsersUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetUsersUnauthorized creates GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

// WithPayload adds the payload to the get users unauthorized response
func (o *GetUsersUnauthorized) WithPayload(payload *models.Unauthorized) *GetUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users unauthorized response
func (o *GetUsersUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUsersForbidden User does not have the credentials to access this resource


swagger:response getUsersForbidden
*/
type GetUsersForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetUsersForbidden creates GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

// WithPayload adds the payload to the get users forbidden response
func (o *GetUsersForbidden) WithPayload(payload *models.Unauthorized) *GetUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users forbidden response
func (o *GetUsersForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUsersNotFound Resource not found

swagger:response getUsersNotFound
*/
type GetUsersNotFound struct {

	// In: body
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetUsersNotFound creates GetUsersNotFound with default headers values
func NewGetUsersNotFound() *GetUsersNotFound {
	return &GetUsersNotFound{}
}

// WithPayload adds the payload to the get users not found response
func (o *GetUsersNotFound) WithPayload(payload *models.NotFound) *GetUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users not found response
func (o *GetUsersNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUsersDefault Error

swagger:response getUsersDefault
*/
type GetUsersDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUsersDefault creates GetUsersDefault with default headers values
func NewGetUsersDefault(code int) *GetUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get users default response
func (o *GetUsersDefault) WithStatusCode(code int) *GetUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get users default response
func (o *GetUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get users default response
func (o *GetUsersDefault) WithPayload(payload *models.Error) *GetUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users default response
func (o *GetUsersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}