// Code generated by go-swagger; DO NOT EDIT.

package salon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SalonCloseHandlerFunc turns a function with the right signature into a salon close handler
type SalonCloseHandlerFunc func(SalonCloseParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SalonCloseHandlerFunc) Handle(params SalonCloseParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SalonCloseHandler interface for that can handle valid salon close params
type SalonCloseHandler interface {
	Handle(SalonCloseParams, interface{}) middleware.Responder
}

// NewSalonClose creates a new http.Handler for the salon close operation
func NewSalonClose(ctx *middleware.Context, handler SalonCloseHandler) *SalonClose {
	return &SalonClose{Context: ctx, Handler: handler}
}

/* SalonClose swagger:route POST /salon/close Salon salonClose

Close salon shutter

Close salon shutter

*/
type SalonClose struct {
	Context *middleware.Context
	Handler SalonCloseHandler
}

func (o *SalonClose) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSalonCloseParams()
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
