// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMessagingServiceMessagesExceptionGetParams creates a new MessagingServiceMessagesExceptionGetParams object
// with the default values initialized.
func NewMessagingServiceMessagesExceptionGetParams() *MessagingServiceMessagesExceptionGetParams {

	return &MessagingServiceMessagesExceptionGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMessagingServiceMessagesExceptionGetParamsWithTimeout creates a new MessagingServiceMessagesExceptionGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMessagingServiceMessagesExceptionGetParamsWithTimeout(timeout time.Duration) *MessagingServiceMessagesExceptionGetParams {

	return &MessagingServiceMessagesExceptionGetParams{

		timeout: timeout,
	}
}

// NewMessagingServiceMessagesExceptionGetParamsWithContext creates a new MessagingServiceMessagesExceptionGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMessagingServiceMessagesExceptionGetParamsWithContext(ctx context.Context) *MessagingServiceMessagesExceptionGetParams {

	return &MessagingServiceMessagesExceptionGetParams{

		Context: ctx,
	}
}

// NewMessagingServiceMessagesExceptionGetParamsWithHTTPClient creates a new MessagingServiceMessagesExceptionGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMessagingServiceMessagesExceptionGetParamsWithHTTPClient(client *http.Client) *MessagingServiceMessagesExceptionGetParams {

	return &MessagingServiceMessagesExceptionGetParams{
		HTTPClient: client,
	}
}

/*
MessagingServiceMessagesExceptionGetParams contains all the parameters to send to the API endpoint
for the messaging service messages exception get operation typically these are written to a http.Request
*/
type MessagingServiceMessagesExceptionGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the messaging service messages exception get params
func (o *MessagingServiceMessagesExceptionGetParams) WithTimeout(timeout time.Duration) *MessagingServiceMessagesExceptionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messaging service messages exception get params
func (o *MessagingServiceMessagesExceptionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messaging service messages exception get params
func (o *MessagingServiceMessagesExceptionGetParams) WithContext(ctx context.Context) *MessagingServiceMessagesExceptionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messaging service messages exception get params
func (o *MessagingServiceMessagesExceptionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messaging service messages exception get params
func (o *MessagingServiceMessagesExceptionGetParams) WithHTTPClient(client *http.Client) *MessagingServiceMessagesExceptionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messaging service messages exception get params
func (o *MessagingServiceMessagesExceptionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MessagingServiceMessagesExceptionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
