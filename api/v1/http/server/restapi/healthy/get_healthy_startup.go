// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package healthy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetHealthyStartupHandlerFunc turns a function with the right signature into a get healthy startup handler
type GetHealthyStartupHandlerFunc func(GetHealthyStartupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHealthyStartupHandlerFunc) Handle(params GetHealthyStartupParams) middleware.Responder {
	return fn(params)
}

// GetHealthyStartupHandler interface for that can handle valid get healthy startup params
type GetHealthyStartupHandler interface {
	Handle(GetHealthyStartupParams) middleware.Responder
}

// NewGetHealthyStartup creates a new http.Handler for the get healthy startup operation
func NewGetHealthyStartup(ctx *middleware.Context, handler GetHealthyStartupHandler) *GetHealthyStartup {
	return &GetHealthyStartup{Context: ctx, Handler: handler}
}

/*
	GetHealthyStartup swagger:route GET /healthy/startup healthy getHealthyStartup

# Startup probe

pod startup probe for agent and controller pod
*/
type GetHealthyStartup struct {
	Context *middleware.Context
	Handler GetHealthyStartupHandler
}

func (o *GetHealthyStartup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHealthyStartupParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
