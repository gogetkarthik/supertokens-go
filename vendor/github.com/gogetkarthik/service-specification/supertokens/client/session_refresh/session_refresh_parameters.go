// Code generated by go-swagger; DO NOT EDIT.

package session_refresh

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

	"github.com/gogetkarthik/service-specification/supertokens/models"
)

// NewSessionRefreshParams creates a new SessionRefreshParams object
// with the default values initialized.
func NewSessionRefreshParams() *SessionRefreshParams {
	var ()
	return &SessionRefreshParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSessionRefreshParamsWithTimeout creates a new SessionRefreshParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSessionRefreshParamsWithTimeout(timeout time.Duration) *SessionRefreshParams {
	var ()
	return &SessionRefreshParams{

		timeout: timeout,
	}
}

// NewSessionRefreshParamsWithContext creates a new SessionRefreshParams object
// with the default values initialized, and the ability to set a context for a request
func NewSessionRefreshParamsWithContext(ctx context.Context) *SessionRefreshParams {
	var ()
	return &SessionRefreshParams{

		Context: ctx,
	}
}

// NewSessionRefreshParamsWithHTTPClient creates a new SessionRefreshParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSessionRefreshParamsWithHTTPClient(client *http.Client) *SessionRefreshParams {
	var ()
	return &SessionRefreshParams{
		HTTPClient: client,
	}
}

/*SessionRefreshParams contains all the parameters to send to the API endpoint
for the session refresh operation typically these are written to a http.Request
*/
type SessionRefreshParams struct {

	/*RefreshSessionInput*/
	RefreshSessionInput *models.RefreshSessionInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the session refresh params
func (o *SessionRefreshParams) WithTimeout(timeout time.Duration) *SessionRefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session refresh params
func (o *SessionRefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session refresh params
func (o *SessionRefreshParams) WithContext(ctx context.Context) *SessionRefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session refresh params
func (o *SessionRefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session refresh params
func (o *SessionRefreshParams) WithHTTPClient(client *http.Client) *SessionRefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session refresh params
func (o *SessionRefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefreshSessionInput adds the refreshSessionInput to the session refresh params
func (o *SessionRefreshParams) WithRefreshSessionInput(refreshSessionInput *models.RefreshSessionInput) *SessionRefreshParams {
	o.SetRefreshSessionInput(refreshSessionInput)
	return o
}

// SetRefreshSessionInput adds the refreshSessionInput to the session refresh params
func (o *SessionRefreshParams) SetRefreshSessionInput(refreshSessionInput *models.RefreshSessionInput) {
	o.RefreshSessionInput = refreshSessionInput
}

// WriteToRequest writes these params to a swagger request
func (o *SessionRefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RefreshSessionInput != nil {
		if err := r.SetBodyParam(o.RefreshSessionInput); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
