//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaActionsPropertiesAddHandlerFunc turns a function with the right signature into a schema actions properties add handler
type SchemaActionsPropertiesAddHandlerFunc func(SchemaActionsPropertiesAddParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SchemaActionsPropertiesAddHandlerFunc) Handle(params SchemaActionsPropertiesAddParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SchemaActionsPropertiesAddHandler interface for that can handle valid schema actions properties add params
type SchemaActionsPropertiesAddHandler interface {
	Handle(SchemaActionsPropertiesAddParams, *models.Principal) middleware.Responder
}

// NewSchemaActionsPropertiesAdd creates a new http.Handler for the schema actions properties add operation
func NewSchemaActionsPropertiesAdd(ctx *middleware.Context, handler SchemaActionsPropertiesAddHandler) *SchemaActionsPropertiesAdd {
	return &SchemaActionsPropertiesAdd{Context: ctx, Handler: handler}
}

/*SchemaActionsPropertiesAdd swagger:route POST /schema/actions/{className}/properties schema schemaActionsPropertiesAdd

Add a property to an Action class.

*/
type SchemaActionsPropertiesAdd struct {
	Context *middleware.Context
	Handler SchemaActionsPropertiesAddHandler
}

func (o *SchemaActionsPropertiesAdd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSchemaActionsPropertiesAddParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}