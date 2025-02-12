// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetFqdnCacheIDReader is a Reader for the GetFqdnCacheID structure.
type GetFqdnCacheIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFqdnCacheIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFqdnCacheIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFqdnCacheIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFqdnCacheIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /fqdn/cache/{id}] GetFqdnCacheID", response, response.Code())
	}
}

// NewGetFqdnCacheIDOK creates a GetFqdnCacheIDOK with default headers values
func NewGetFqdnCacheIDOK() *GetFqdnCacheIDOK {
	return &GetFqdnCacheIDOK{}
}

/*
GetFqdnCacheIDOK describes a response with status code 200, with default header values.

Success
*/
type GetFqdnCacheIDOK struct {
	Payload []*models.DNSLookup
}

// IsSuccess returns true when this get fqdn cache Id o k response has a 2xx status code
func (o *GetFqdnCacheIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fqdn cache Id o k response has a 3xx status code
func (o *GetFqdnCacheIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fqdn cache Id o k response has a 4xx status code
func (o *GetFqdnCacheIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fqdn cache Id o k response has a 5xx status code
func (o *GetFqdnCacheIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fqdn cache Id o k response a status code equal to that given
func (o *GetFqdnCacheIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get fqdn cache Id o k response
func (o *GetFqdnCacheIDOK) Code() int {
	return 200
}

func (o *GetFqdnCacheIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdOK %s", 200, payload)
}

func (o *GetFqdnCacheIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdOK %s", 200, payload)
}

func (o *GetFqdnCacheIDOK) GetPayload() []*models.DNSLookup {
	return o.Payload
}

func (o *GetFqdnCacheIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFqdnCacheIDBadRequest creates a GetFqdnCacheIDBadRequest with default headers values
func NewGetFqdnCacheIDBadRequest() *GetFqdnCacheIDBadRequest {
	return &GetFqdnCacheIDBadRequest{}
}

/*
GetFqdnCacheIDBadRequest describes a response with status code 400, with default header values.

Invalid request (error parsing parameters)
*/
type GetFqdnCacheIDBadRequest struct {
	Payload models.Error
}

// IsSuccess returns true when this get fqdn cache Id bad request response has a 2xx status code
func (o *GetFqdnCacheIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fqdn cache Id bad request response has a 3xx status code
func (o *GetFqdnCacheIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fqdn cache Id bad request response has a 4xx status code
func (o *GetFqdnCacheIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fqdn cache Id bad request response has a 5xx status code
func (o *GetFqdnCacheIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fqdn cache Id bad request response a status code equal to that given
func (o *GetFqdnCacheIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get fqdn cache Id bad request response
func (o *GetFqdnCacheIDBadRequest) Code() int {
	return 400
}

func (o *GetFqdnCacheIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdBadRequest %s", 400, payload)
}

func (o *GetFqdnCacheIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdBadRequest %s", 400, payload)
}

func (o *GetFqdnCacheIDBadRequest) GetPayload() models.Error {
	return o.Payload
}

func (o *GetFqdnCacheIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFqdnCacheIDNotFound creates a GetFqdnCacheIDNotFound with default headers values
func NewGetFqdnCacheIDNotFound() *GetFqdnCacheIDNotFound {
	return &GetFqdnCacheIDNotFound{}
}

/*
GetFqdnCacheIDNotFound describes a response with status code 404, with default header values.

No DNS data with provided parameters found
*/
type GetFqdnCacheIDNotFound struct {
}

// IsSuccess returns true when this get fqdn cache Id not found response has a 2xx status code
func (o *GetFqdnCacheIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fqdn cache Id not found response has a 3xx status code
func (o *GetFqdnCacheIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fqdn cache Id not found response has a 4xx status code
func (o *GetFqdnCacheIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fqdn cache Id not found response has a 5xx status code
func (o *GetFqdnCacheIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fqdn cache Id not found response a status code equal to that given
func (o *GetFqdnCacheIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get fqdn cache Id not found response
func (o *GetFqdnCacheIDNotFound) Code() int {
	return 404
}

func (o *GetFqdnCacheIDNotFound) Error() string {
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdNotFound", 404)
}

func (o *GetFqdnCacheIDNotFound) String() string {
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdNotFound", 404)
}

func (o *GetFqdnCacheIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
