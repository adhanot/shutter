// Code generated by go-swagger; DO NOT EDIT.

package sdb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SdbOpenHandlerFunc turns a function with the right signature into a sdb open handler
type SdbOpenHandlerFunc func(SdbOpenParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SdbOpenHandlerFunc) Handle(params SdbOpenParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SdbOpenHandler interface for that can handle valid sdb open params
type SdbOpenHandler interface {
	Handle(SdbOpenParams, interface{}) middleware.Responder
}

// NewSdbOpen creates a new http.Handler for the sdb open operation
func NewSdbOpen(ctx *middleware.Context, handler SdbOpenHandler) *SdbOpen {
	return &SdbOpen{Context: ctx, Handler: handler}
}

/* SdbOpen swagger:route POST /sdb/open Sdb sdbOpen

Open sdb shutter

Open sdb shutter

*/
type SdbOpen struct {
	Context *middleware.Context
	Handler SdbOpenHandler
}

func (o *SdbOpen) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSdbOpenParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}