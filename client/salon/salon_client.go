// Code generated by go-swagger; DO NOT EDIT.

package salon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new salon API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for salon API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SalonClose(params *SalonCloseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SalonCloseOK, error)

	SalonOpen(params *SalonOpenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SalonOpenOK, error)

	SalonStop(params *SalonStopParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SalonStopOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SalonClose closes salon shutter

  Close salon shutter
*/
func (a *Client) SalonClose(params *SalonCloseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SalonCloseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalonCloseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Salon/close",
		Method:             "POST",
		PathPattern:        "/salon/close",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalonCloseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalonCloseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Salon/close: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalonOpen opens salon shutter

  Open salon shutter
*/
func (a *Client) SalonOpen(params *SalonOpenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SalonOpenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalonOpenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Salon/open",
		Method:             "POST",
		PathPattern:        "/salon/open",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalonOpenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalonOpenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Salon/open: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalonStop stops salon shutter

  Stop salon shutter
*/
func (a *Client) SalonStop(params *SalonStopParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SalonStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalonStopParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Salon/stop",
		Method:             "POST",
		PathPattern:        "/salon/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalonStopReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalonStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Salon/stop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
