// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigReduceCacheCapacityToParams creates a new FindConfigReduceCacheCapacityToParams object
// with the default values initialized.
func NewFindConfigReduceCacheCapacityToParams() *FindConfigReduceCacheCapacityToParams {

	return &FindConfigReduceCacheCapacityToParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigReduceCacheCapacityToParamsWithTimeout creates a new FindConfigReduceCacheCapacityToParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigReduceCacheCapacityToParamsWithTimeout(timeout time.Duration) *FindConfigReduceCacheCapacityToParams {

	return &FindConfigReduceCacheCapacityToParams{

		timeout: timeout,
	}
}

// NewFindConfigReduceCacheCapacityToParamsWithContext creates a new FindConfigReduceCacheCapacityToParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigReduceCacheCapacityToParamsWithContext(ctx context.Context) *FindConfigReduceCacheCapacityToParams {

	return &FindConfigReduceCacheCapacityToParams{

		Context: ctx,
	}
}

// NewFindConfigReduceCacheCapacityToParamsWithHTTPClient creates a new FindConfigReduceCacheCapacityToParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigReduceCacheCapacityToParamsWithHTTPClient(client *http.Client) *FindConfigReduceCacheCapacityToParams {

	return &FindConfigReduceCacheCapacityToParams{
		HTTPClient: client,
	}
}

/*
FindConfigReduceCacheCapacityToParams contains all the parameters to send to the API endpoint
for the find config reduce cache capacity to operation typically these are written to a http.Request
*/
type FindConfigReduceCacheCapacityToParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config reduce cache capacity to params
func (o *FindConfigReduceCacheCapacityToParams) WithTimeout(timeout time.Duration) *FindConfigReduceCacheCapacityToParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config reduce cache capacity to params
func (o *FindConfigReduceCacheCapacityToParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config reduce cache capacity to params
func (o *FindConfigReduceCacheCapacityToParams) WithContext(ctx context.Context) *FindConfigReduceCacheCapacityToParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config reduce cache capacity to params
func (o *FindConfigReduceCacheCapacityToParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config reduce cache capacity to params
func (o *FindConfigReduceCacheCapacityToParams) WithHTTPClient(client *http.Client) *FindConfigReduceCacheCapacityToParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config reduce cache capacity to params
func (o *FindConfigReduceCacheCapacityToParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigReduceCacheCapacityToParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
