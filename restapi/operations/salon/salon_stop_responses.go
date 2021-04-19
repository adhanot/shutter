// Code generated by go-swagger; DO NOT EDIT.

package salon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"shutter/models"
)

// SalonStopOKCode is the HTTP code returned for type SalonStopOK
const SalonStopOKCode int = 200

/*SalonStopOK Successful Stop

swagger:response salonStopOK
*/
type SalonStopOK struct {

	/*
	  In: Body
	*/
	Payload *models.Success `json:"body,omitempty"`
}

// NewSalonStopOK creates SalonStopOK with default headers values
func NewSalonStopOK() *SalonStopOK {

	return &SalonStopOK{}
}

// WithPayload adds the payload to the salon stop o k response
func (o *SalonStopOK) WithPayload(payload *models.Success) *SalonStopOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the salon stop o k response
func (o *SalonStopOK) SetPayload(payload *models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalonStopOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SalonStopBadRequestCode is the HTTP code returned for type SalonStopBadRequest
const SalonStopBadRequestCode int = 400

/*SalonStopBadRequest Error

swagger:response salonStopBadRequest
*/
type SalonStopBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSalonStopBadRequest creates SalonStopBadRequest with default headers values
func NewSalonStopBadRequest() *SalonStopBadRequest {

	return &SalonStopBadRequest{}
}

// WithPayload adds the payload to the salon stop bad request response
func (o *SalonStopBadRequest) WithPayload(payload *models.Error) *SalonStopBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the salon stop bad request response
func (o *SalonStopBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalonStopBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SalonStopInternalServerErrorCode is the HTTP code returned for type SalonStopInternalServerError
const SalonStopInternalServerErrorCode int = 500

/*SalonStopInternalServerError Server error

swagger:response salonStopInternalServerError
*/
type SalonStopInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewSalonStopInternalServerError creates SalonStopInternalServerError with default headers values
func NewSalonStopInternalServerError() *SalonStopInternalServerError {

	return &SalonStopInternalServerError{}
}

// WithPayload adds the payload to the salon stop internal server error response
func (o *SalonStopInternalServerError) WithPayload(payload *models.ServerError) *SalonStopInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the salon stop internal server error response
func (o *SalonStopInternalServerError) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalonStopInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}