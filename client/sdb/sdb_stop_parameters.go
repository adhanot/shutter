// Code generated by go-swagger; DO NOT EDIT.

package sdb

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

// NewSdbStopParams creates a new SdbStopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSdbStopParams() *SdbStopParams {
	return &SdbStopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSdbStopParamsWithTimeout creates a new SdbStopParams object
// with the ability to set a timeout on a request.
func NewSdbStopParamsWithTimeout(timeout time.Duration) *SdbStopParams {
	return &SdbStopParams{
		timeout: timeout,
	}
}

// NewSdbStopParamsWithContext creates a new SdbStopParams object
// with the ability to set a context for a request.
func NewSdbStopParamsWithContext(ctx context.Context) *SdbStopParams {
	return &SdbStopParams{
		Context: ctx,
	}
}

// NewSdbStopParamsWithHTTPClient creates a new SdbStopParams object
// with the ability to set a custom HTTPClient for a request.
func NewSdbStopParamsWithHTTPClient(client *http.Client) *SdbStopParams {
	return &SdbStopParams{
		HTTPClient: client,
	}
}

/* SdbStopParams contains all the parameters to send to the API endpoint
   for the sdb stop operation.

   Typically these are written to a http.Request.
*/
type SdbStopParams struct {

	/* Stop.

	   Registeration Payload
	*/
	Stop models.Shutter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sdb stop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SdbStopParams) WithDefaults() *SdbStopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sdb stop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SdbStopParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sdb stop params
func (o *SdbStopParams) WithTimeout(timeout time.Duration) *SdbStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sdb stop params
func (o *SdbStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sdb stop params
func (o *SdbStopParams) WithContext(ctx context.Context) *SdbStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sdb stop params
func (o *SdbStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sdb stop params
func (o *SdbStopParams) WithHTTPClient(client *http.Client) *SdbStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sdb stop params
func (o *SdbStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStop adds the stop to the sdb stop params
func (o *SdbStopParams) WithStop(stop models.Shutter) *SdbStopParams {
	o.SetStop(stop)
	return o
}

// SetStop adds the stop to the sdb stop params
func (o *SdbStopParams) SetStop(stop models.Shutter) {
	o.Stop = stop
}

// WriteToRequest writes these params to a swagger request
func (o *SdbStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Stop != nil {
		if err := r.SetBodyParam(o.Stop); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
