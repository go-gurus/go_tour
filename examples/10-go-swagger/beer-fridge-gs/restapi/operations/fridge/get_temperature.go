// Code generated by go-swagger; DO NOT EDIT.

package fridge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTemperatureHandlerFunc turns a function with the right signature into a get temperature handler
type GetTemperatureHandlerFunc func(GetTemperatureParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTemperatureHandlerFunc) Handle(params GetTemperatureParams) middleware.Responder {
	return fn(params)
}

// GetTemperatureHandler interface for that can handle valid get temperature params
type GetTemperatureHandler interface {
	Handle(GetTemperatureParams) middleware.Responder
}

// NewGetTemperature creates a new http.Handler for the get temperature operation
func NewGetTemperature(ctx *middleware.Context, handler GetTemperatureHandler) *GetTemperature {
	return &GetTemperature{Context: ctx, Handler: handler}
}

/* GetTemperature swagger:route GET /temperature fridge getTemperature

GetTemperature get temperature API

*/
type GetTemperature struct {
	Context *middleware.Context
	Handler GetTemperatureHandler
}

func (o *GetTemperature) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTemperatureParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
