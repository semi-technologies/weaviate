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

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// C11yCorpusGetHandlerFunc turns a function with the right signature into a c11y corpus get handler
type C11yCorpusGetHandlerFunc func(C11yCorpusGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn C11yCorpusGetHandlerFunc) Handle(params C11yCorpusGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// C11yCorpusGetHandler interface for that can handle valid c11y corpus get params
type C11yCorpusGetHandler interface {
	Handle(C11yCorpusGetParams, *models.Principal) middleware.Responder
}

// NewC11yCorpusGet creates a new http.Handler for the c11y corpus get operation
func NewC11yCorpusGet(ctx *middleware.Context, handler C11yCorpusGetHandler) *C11yCorpusGet {
	return &C11yCorpusGet{Context: ctx, Handler: handler}
}

/*C11yCorpusGet swagger:route POST /c11y/corpus contextionary-API c11yCorpusGet

Checks if a word or wordString is part of the contextionary.

Analyzes a sentence based on the contextionary

*/
type C11yCorpusGet struct {
	Context *middleware.Context
	Handler C11yCorpusGetHandler
}

func (o *C11yCorpusGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewC11yCorpusGetParams()

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

// C11yCorpusGetBody The text corpus.
// swagger:model C11yCorpusGetBody
type C11yCorpusGetBody struct {

	// corpus
	Corpus string `json:"corpus,omitempty"`
}

// Validate validates this c11y corpus get body
func (o *C11yCorpusGetBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *C11yCorpusGetBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *C11yCorpusGetBody) UnmarshalBinary(b []byte) error {
	var res C11yCorpusGetBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}