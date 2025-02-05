// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigBroadcastRPCAddressReader is a Reader for the FindConfigBroadcastRPCAddress structure.
type FindConfigBroadcastRPCAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigBroadcastRPCAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigBroadcastRPCAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigBroadcastRPCAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigBroadcastRPCAddressOK creates a FindConfigBroadcastRPCAddressOK with default headers values
func NewFindConfigBroadcastRPCAddressOK() *FindConfigBroadcastRPCAddressOK {
	return &FindConfigBroadcastRPCAddressOK{}
}

/*
FindConfigBroadcastRPCAddressOK handles this case with default header values.

Config value
*/
type FindConfigBroadcastRPCAddressOK struct {
	Payload string
}

func (o *FindConfigBroadcastRPCAddressOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigBroadcastRPCAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigBroadcastRPCAddressDefault creates a FindConfigBroadcastRPCAddressDefault with default headers values
func NewFindConfigBroadcastRPCAddressDefault(code int) *FindConfigBroadcastRPCAddressDefault {
	return &FindConfigBroadcastRPCAddressDefault{
		_statusCode: code,
	}
}

/*
FindConfigBroadcastRPCAddressDefault handles this case with default header values.

unexpected error
*/
type FindConfigBroadcastRPCAddressDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config broadcast rpc address default response
func (o *FindConfigBroadcastRPCAddressDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigBroadcastRPCAddressDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigBroadcastRPCAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigBroadcastRPCAddressDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
