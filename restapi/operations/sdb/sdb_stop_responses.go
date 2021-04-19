// Code generated by go-swagger; DO NOT EDIT.

package sdb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"shutter/models"
)

// SdbStopOKCode is the HTTP code returned for type SdbStopOK
const SdbStopOKCode int = 200

/*SdbStopOK Successful Stop

swagger:response sdbStopOK
*/
type SdbStopOK struct {

	/*
	  In: Body
	*/
	Payload *models.Success `json:"body,omitempty"`
}

// NewSdbStopOK creates SdbStopOK with default headers values
func NewSdbStopOK() *SdbStopOK {

	return &SdbStopOK{}
}

// WithPayload adds the payload to the sdb stop o k response
func (o *SdbStopOK) WithPayload(payload *models.Success) *SdbStopOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sdb stop o k response
func (o *SdbStopOK) SetPayload(payload *models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SdbStopOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SdbStopBadRequestCode is the HTTP code returned for type SdbStopBadRequest
const SdbStopBadRequestCode int = 400

/*SdbStopBadRequest Error

swagger:response sdbStopBadRequest
*/
type SdbStopBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSdbStopBadRequest creates SdbStopBadRequest with default headers values
func NewSdbStopBadRequest() *SdbStopBadRequest {

	return &SdbStopBadRequest{}
}

// WithPayload adds the payload to the sdb stop bad request response
func (o *SdbStopBadRequest) WithPayload(payload *models.Error) *SdbStopBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sdb stop bad request response
func (o *SdbStopBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SdbStopBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SdbStopInternalServerErrorCode is the HTTP code returned for type SdbStopInternalServerError
const SdbStopInternalServerErrorCode int = 500

/*SdbStopInternalServerError Server error

swagger:response sdbStopInternalServerError
*/
type SdbStopInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewSdbStopInternalServerError creates SdbStopInternalServerError with default headers values
func NewSdbStopInternalServerError() *SdbStopInternalServerError {

	return &SdbStopInternalServerError{}
}

// WithPayload adds the payload to the sdb stop internal server error response
func (o *SdbStopInternalServerError) WithPayload(payload *models.ServerError) *SdbStopInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sdb stop internal server error response
func (o *SdbStopInternalServerError) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SdbStopInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}