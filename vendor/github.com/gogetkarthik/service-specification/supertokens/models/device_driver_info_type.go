// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceDriverInfoType device driver info type
//
// swagger:model DeviceDriverInfoType
type DeviceDriverInfoType struct {

	// driver
	Driver *Driver `json:"driver,omitempty"`

	// frontend s d k
	FrontendSDK []*FrontendSDK `json:"frontendSDK"`
}

// Validate validates this device driver info type
func (m *DeviceDriverInfoType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontendSDK(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceDriverInfoType) validateDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.Driver) { // not required
		return nil
	}

	if m.Driver != nil {
		if err := m.Driver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceDriverInfoType) validateFrontendSDK(formats strfmt.Registry) error {

	if swag.IsZero(m.FrontendSDK) { // not required
		return nil
	}

	for i := 0; i < len(m.FrontendSDK); i++ {
		if swag.IsZero(m.FrontendSDK[i]) { // not required
			continue
		}

		if m.FrontendSDK[i] != nil {
			if err := m.FrontendSDK[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frontendSDK" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceDriverInfoType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceDriverInfoType) UnmarshalBinary(b []byte) error {
	var res DeviceDriverInfoType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
