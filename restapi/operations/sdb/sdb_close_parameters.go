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

// NewSdbCloseParams creates a new SdbCloseParams object
//
// There are no default values defined in the spec.
func NewSdbCloseParams() SdbCloseParams {

	return SdbCloseParams{}
}

// SdbCloseParams contains all the bound params for the sdb close operation
// typically these are obtained from a http.Request
//
// swagger:parameters sdb/close
type SdbCloseParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Registeration Payload
	  In: body
	*/
	Close models.Shutter
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSdbCloseParams() beforehand.
func (o *SdbCloseParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Shutter
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("close", "body", "", err))
		} else {
			// no validation on generic interface
			o.Close = body
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
