// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Token Token details
//
// swagger:model Token
type Token struct {

	// cookie path
	CookiePath string `json:"cookiePath,omitempty"`

	// cookie secure
	CookieSecure bool `json:"cookieSecure,omitempty"`

	// created time
	CreatedTime int64 `json:"createdTime,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// expiry
	Expiry int64 `json:"expiry,omitempty"`

	// same site
	SameSite string `json:"sameSite,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this token
func (m *Token) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Token) UnmarshalBinary(b []byte) error {
	var res Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
