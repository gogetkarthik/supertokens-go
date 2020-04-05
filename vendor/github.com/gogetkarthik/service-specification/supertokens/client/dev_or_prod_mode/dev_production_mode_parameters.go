// Code generated by go-swagger; DO NOT EDIT.

package dev_or_prod_mode

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

// NewDevProductionModeParams creates a new DevProductionModeParams object
// with the default values initialized.
func NewDevProductionModeParams() *DevProductionModeParams {
	var ()
	return &DevProductionModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevProductionModeParamsWithTimeout creates a new DevProductionModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevProductionModeParamsWithTimeout(timeout time.Duration) *DevProductionModeParams {
	var ()
	return &DevProductionModeParams{

		timeout: timeout,
	}
}

// NewDevProductionModeParamsWithContext creates a new DevProductionModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevProductionModeParamsWithContext(ctx context.Context) *DevProductionModeParams {
	var ()
	return &DevProductionModeParams{

		Context: ctx,
	}
}

// NewDevProductionModeParamsWithHTTPClient creates a new DevProductionModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevProductionModeParamsWithHTTPClient(client *http.Client) *DevProductionModeParams {
	var ()
	return &DevProductionModeParams{
		HTTPClient: client,
	}
}

/*DevProductionModeParams contains all the parameters to send to the API endpoint
for the dev production mode operation typically these are written to a http.Request
*/
type DevProductionModeParams struct {

	/*Pid*/
	Pid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dev production mode params
func (o *DevProductionModeParams) WithTimeout(timeout time.Duration) *DevProductionModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dev production mode params
func (o *DevProductionModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dev production mode params
func (o *DevProductionModeParams) WithContext(ctx context.Context) *DevProductionModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dev production mode params
func (o *DevProductionModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dev production mode params
func (o *DevProductionModeParams) WithHTTPClient(client *http.Client) *DevProductionModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dev production mode params
func (o *DevProductionModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPid adds the pid to the dev production mode params
func (o *DevProductionModeParams) WithPid(pid string) *DevProductionModeParams {
	o.SetPid(pid)
	return o
}

// SetPid adds the pid to the dev production mode params
func (o *DevProductionModeParams) SetPid(pid string) {
	o.Pid = pid
}

// WriteToRequest writes these params to a swagger request
func (o *DevProductionModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param pid
	qrPid := o.Pid
	qPid := qrPid
	if qPid != "" {
		if err := r.SetQueryParam("pid", qPid); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
