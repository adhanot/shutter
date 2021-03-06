// Code generated by go-swagger; DO NOT EDIT.

package salon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"shutter/models"
)

// NewSalonOpenParams creates a new SalonOpenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSalonOpenParams() *SalonOpenParams {
	return &SalonOpenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSalonOpenParamsWithTimeout creates a new SalonOpenParams object
// with the ability to set a timeout on a request.
func NewSalonOpenParamsWithTimeout(timeout time.Duration) *SalonOpenParams {
	return &SalonOpenParams{
		timeout: timeout,
	}
}

// NewSalonOpenParamsWithContext creates a new SalonOpenParams object
// with the ability to set a context for a request.
func NewSalonOpenParamsWithContext(ctx context.Context) *SalonOpenParams {
	return &SalonOpenParams{
		Context: ctx,
	}
}

// NewSalonOpenParamsWithHTTPClient creates a new SalonOpenParams object
// with the ability to set a custom HTTPClient for a request.
func NewSalonOpenParamsWithHTTPClient(client *http.Client) *SalonOpenParams {
	return &SalonOpenParams{
		HTTPClient: client,
	}
}

/* SalonOpenParams contains all the parameters to send to the API endpoint
   for the salon open operation.

   Typically these are written to a http.Request.
*/
type SalonOpenParams struct {

	/* Open.

	   Registeration Payload
	*/
	Open models.Shutter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the salon open params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalonOpenParams) WithDefaults() *SalonOpenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the salon open params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalonOpenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the salon open params
func (o *SalonOpenParams) WithTimeout(timeout time.Duration) *SalonOpenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the salon open params
func (o *SalonOpenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the salon open params
func (o *SalonOpenParams) WithContext(ctx context.Context) *SalonOpenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the salon open params
func (o *SalonOpenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the salon open params
func (o *SalonOpenParams) WithHTTPClient(client *http.Client) *SalonOpenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the salon open params
func (o *SalonOpenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpen adds the open to the salon open params
func (o *SalonOpenParams) WithOpen(open models.Shutter) *SalonOpenParams {
	o.SetOpen(open)
	return o
}

// SetOpen adds the open to the salon open params
func (o *SalonOpenParams) SetOpen(open models.Shutter) {
	o.Open = open
}

// WriteToRequest writes these params to a swagger request
func (o *SalonOpenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Open != nil {
		if err := r.SetBodyParam(o.Open); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
