// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHelloGetParams creates a new HelloGetParams object
// with the default values initialized.
func NewHelloGetParams() *HelloGetParams {

	return &HelloGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHelloGetParamsWithTimeout creates a new HelloGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHelloGetParamsWithTimeout(timeout time.Duration) *HelloGetParams {

	return &HelloGetParams{

		timeout: timeout,
	}
}

// NewHelloGetParamsWithContext creates a new HelloGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewHelloGetParamsWithContext(ctx context.Context) *HelloGetParams {

	return &HelloGetParams{

		Context: ctx,
	}
}

// NewHelloGetParamsWithHTTPClient creates a new HelloGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHelloGetParamsWithHTTPClient(client *http.Client) *HelloGetParams {

	return &HelloGetParams{
		HTTPClient: client,
	}
}

/*HelloGetParams contains all the parameters to send to the API endpoint
for the hello get operation typically these are written to a http.Request
*/
type HelloGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hello get params
func (o *HelloGetParams) WithTimeout(timeout time.Duration) *HelloGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hello get params
func (o *HelloGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hello get params
func (o *HelloGetParams) WithContext(ctx context.Context) *HelloGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hello get params
func (o *HelloGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hello get params
func (o *HelloGetParams) WithHTTPClient(client *http.Client) *HelloGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hello get params
func (o *HelloGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HelloGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
