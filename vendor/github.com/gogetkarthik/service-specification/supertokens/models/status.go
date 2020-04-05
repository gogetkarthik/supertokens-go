// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Status status
//
// swagger:model Status
type Status string

const (

	// StatusOK captures enum value "OK"
	StatusOK Status = "OK"

	// StatusNOTALLOWED captures enum value "NOT_ALLOWED"
	StatusNOTALLOWED Status = "NOT_ALLOWED"

	// StatusNOTSUPPORTED captures enum value "NOT_SUPPORTED"
	StatusNOTSUPPORTED Status = "NOT_SUPPORTED"

	// StatusTOKENTHEFTDETECTED captures enum value "TOKEN_THEFT_DETECTED"
	StatusTOKENTHEFTDETECTED Status = "TOKEN_THEFT_DETECTED"

	// StatusTRYREFRESHTOKEN captures enum value "TRY_REFRESH_TOKEN"
	StatusTRYREFRESHTOKEN Status = "TRY_REFRESH_TOKEN"

	// StatusUNAUTHORISED captures enum value "UNAUTHORISED"
	StatusUNAUTHORISED Status = "UNAUTHORISED"
)

// for schema
var statusEnum []interface{}

func init() {
	var res []Status
	if err := json.Unmarshal([]byte(`["OK","NOT_ALLOWED","NOT_SUPPORTED","TOKEN_THEFT_DETECTED","TRY_REFRESH_TOKEN","UNAUTHORISED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusEnum = append(statusEnum, v)
	}
}

func (m Status) validateStatusEnum(path, location string, value Status) error {
	if err := validate.Enum(path, location, value, statusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this status
func (m Status) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}