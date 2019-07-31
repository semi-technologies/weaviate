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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsPatchReader is a Reader for the ActionsPatch structure.
type ActionsPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActionsPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewActionsPatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewActionsPatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewActionsPatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewActionsPatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewActionsPatchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewActionsPatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsPatchOK creates a ActionsPatchOK with default headers values
func NewActionsPatchOK() *ActionsPatchOK {
	return &ActionsPatchOK{}
}

/*ActionsPatchOK handles this case with default header values.

Successfully applied.
*/
type ActionsPatchOK struct {
	Payload *models.Action
}

func (o *ActionsPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /actions/{id}][%d] actionsPatchOK  %+v", 200, o.Payload)
}

func (o *ActionsPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsPatchBadRequest creates a ActionsPatchBadRequest with default headers values
func NewActionsPatchBadRequest() *ActionsPatchBadRequest {
	return &ActionsPatchBadRequest{}
}

/*ActionsPatchBadRequest handles this case with default header values.

The patch-JSON is malformed.
*/
type ActionsPatchBadRequest struct {
}

func (o *ActionsPatchBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /actions/{id}][%d] actionsPatchBadRequest ", 400)
}

func (o *ActionsPatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsPatchUnauthorized creates a ActionsPatchUnauthorized with default headers values
func NewActionsPatchUnauthorized() *ActionsPatchUnauthorized {
	return &ActionsPatchUnauthorized{}
}

/*ActionsPatchUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ActionsPatchUnauthorized struct {
}

func (o *ActionsPatchUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /actions/{id}][%d] actionsPatchUnauthorized ", 401)
}

func (o *ActionsPatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsPatchForbidden creates a ActionsPatchForbidden with default headers values
func NewActionsPatchForbidden() *ActionsPatchForbidden {
	return &ActionsPatchForbidden{}
}

/*ActionsPatchForbidden handles this case with default header values.

Forbidden
*/
type ActionsPatchForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ActionsPatchForbidden) Error() string {
	return fmt.Sprintf("[PATCH /actions/{id}][%d] actionsPatchForbidden  %+v", 403, o.Payload)
}

func (o *ActionsPatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsPatchNotFound creates a ActionsPatchNotFound with default headers values
func NewActionsPatchNotFound() *ActionsPatchNotFound {
	return &ActionsPatchNotFound{}
}

/*ActionsPatchNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ActionsPatchNotFound struct {
}

func (o *ActionsPatchNotFound) Error() string {
	return fmt.Sprintf("[PATCH /actions/{id}][%d] actionsPatchNotFound ", 404)
}

func (o *ActionsPatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsPatchUnprocessableEntity creates a ActionsPatchUnprocessableEntity with default headers values
func NewActionsPatchUnprocessableEntity() *ActionsPatchUnprocessableEntity {
	return &ActionsPatchUnprocessableEntity{}
}

/*ActionsPatchUnprocessableEntity handles this case with default header values.

The patch-JSON is valid but unprocessable.
*/
type ActionsPatchUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ActionsPatchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /actions/{id}][%d] actionsPatchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ActionsPatchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsPatchInternalServerError creates a ActionsPatchInternalServerError with default headers values
func NewActionsPatchInternalServerError() *ActionsPatchInternalServerError {
	return &ActionsPatchInternalServerError{}
}

/*ActionsPatchInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ActionsPatchInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ActionsPatchInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /actions/{id}][%d] actionsPatchInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionsPatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}