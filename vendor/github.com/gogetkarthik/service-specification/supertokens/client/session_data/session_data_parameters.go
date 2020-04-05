// Code generated by go-swagger; DO NOT EDIT.

package session_data

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

// NewSessionDataParams creates a new SessionDataParams object
// with the default values initialized.
func NewSessionDataParams() *SessionDataParams {
	var ()
	return &SessionDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSessionDataParamsWithTimeout creates a new SessionDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSessionDataParamsWithTimeout(timeout time.Duration) *SessionDataParams {
	var ()
	return &SessionDataParams{

		timeout: timeout,
	}
}

// NewSessionDataParamsWithContext creates a new SessionDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewSessionDataParamsWithContext(ctx context.Context) *SessionDataParams {
	var ()
	return &SessionDataParams{

		Context: ctx,
	}
}

// NewSessionDataParamsWithHTTPClient creates a new SessionDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSessionDataParamsWithHTTPClient(client *http.Client) *SessionDataParams {
	var ()
	return &SessionDataParams{
		HTTPClient: client,
	}
}

/*SessionDataParams contains all the parameters to send to the API endpoint
for the session data operation typically these are written to a http.Request
*/
type SessionDataParams struct {

	/*SessionHandle*/
	SessionHandle string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the session data params
func (o *SessionDataParams) WithTimeout(timeout time.Duration) *SessionDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session data params
func (o *SessionDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session data params
func (o *SessionDataParams) WithContext(ctx context.Context) *SessionDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session data params
func (o *SessionDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session data params
func (o *SessionDataParams) WithHTTPClient(client *http.Client) *SessionDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session data params
func (o *SessionDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionHandle adds the sessionHandle to the session data params
func (o *SessionDataParams) WithSessionHandle(sessionHandle string) *SessionDataParams {
	o.SetSessionHandle(sessionHandle)
	return o
}

// SetSessionHandle adds the sessionHandle to the session data params
func (o *SessionDataParams) SetSessionHandle(sessionHandle string) {
	o.SessionHandle = sessionHandle
}

// WriteToRequest writes these params to a swagger request
func (o *SessionDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param session_handle
	qrSessionHandle := o.SessionHandle
	qSessionHandle := qrSessionHandle
	if qSessionHandle != "" {
		if err := r.SetQueryParam("session_handle", qSessionHandle); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}