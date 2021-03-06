// Code generated by go-swagger; DO NOT EDIT.

package salon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"shutter/models"
)

// SalonCloseOKCode is the HTTP code returned for type SalonCloseOK
const SalonCloseOKCode int = 200

/*SalonCloseOK Successful Close

swagger:response salonCloseOK
*/
type SalonCloseOK struct {

	/*
	  In: Body
	*/
	Payload *models.Success `json:"body,omitempty"`
}

// NewSalonCloseOK creates SalonCloseOK with default headers values
func NewSalonCloseOK() *SalonCloseOK {

	return &SalonCloseOK{}
}

// WithPayload adds the payload to the salon close o k response
func (o *SalonCloseOK) WithPayload(payload *models.Success) *SalonCloseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the salon close o k response
func (o *SalonCloseOK) SetPayload(payload *models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalonCloseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SalonCloseBadRequestCode is the HTTP code returned for type SalonCloseBadRequest
const SalonCloseBadRequestCode int = 400

/*SalonCloseBadRequest Error

swagger:response salonCloseBadRequest
*/
type SalonCloseBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSalonCloseBadRequest creates SalonCloseBadRequest with default headers values
func NewSalonCloseBadRequest() *SalonCloseBadRequest {

	return &SalonCloseBadRequest{}
}

// WithPayload adds the payload to the salon close bad request response
func (o *SalonCloseBadRequest) WithPayload(payload *models.Error) *SalonCloseBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the salon close bad request response
func (o *SalonCloseBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalonCloseBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SalonCloseInternalServerErrorCode is the HTTP code returned for type SalonCloseInternalServerError
const SalonCloseInternalServerErrorCode int = 500

/*SalonCloseInternalServerError Server error

swagger:response salonCloseInternalServerError
*/
type SalonCloseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewSalonCloseInternalServerError creates SalonCloseInternalServerError with default headers values
func NewSalonCloseInternalServerError() *SalonCloseInternalServerError {

	return &SalonCloseInternalServerError{}
}

// WithPayload adds the payload to the salon close internal server error response
func (o *SalonCloseInternalServerError) WithPayload(payload *models.ServerError) *SalonCloseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the salon close internal server error response
func (o *SalonCloseInternalServerError) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalonCloseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
