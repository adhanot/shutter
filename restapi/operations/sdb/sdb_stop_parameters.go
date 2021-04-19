// Code generated by go-swagger; DO NOT EDIT.

package sdb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"shutter/models"
)

// NewSdbStopParams creates a new SdbStopParams object
//
// There are no default values defined in the spec.
func NewSdbStopParams() SdbStopParams {

	return SdbStopParams{}
}

// SdbStopParams contains all the bound params for the sdb stop operation
// typically these are obtained from a http.Request
//
// swagger:parameters sdb/stop
type SdbStopParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Registeration Payload
	  In: body
	*/
	Stop models.Shutter
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSdbStopParams() beforehand.
func (o *SdbStopParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Shutter
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("stop", "body", "", err))
		} else {
			// no validation on generic interface
			o.Stop = body
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}